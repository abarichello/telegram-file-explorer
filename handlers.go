package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/mymmrac/telego"

	thandler "github.com/mymmrac/telego/telegohandler"
	tutil "github.com/mymmrac/telego/telegoutil"
)

func setMyCommands() {
	bot.SetMyCommands(&telego.SetMyCommandsParams{
		Commands: []telego.BotCommand{
			{Command: "start", Description: "Starts the bot by listing the root directory"},
			{Command: "help", Description: "Shows how to use"},
		},
	})
}

func registerHandlers() {
	botHandler.HandleMessage(listCommand, thandler.CommandEqual("start"))
	botHandler.HandleCallbackQuery(listCallback, thandler.AnyCallbackQueryWithMessage())
	botHandler.HandleMessage(unknown)
}

func listCallback(bot *telego.Bot, callback telego.CallbackQuery) {
	// ignore first inline keyboard button callback, we need to send data
	// since pure text buttons with no callback are not allowed
	if callback.Data == "empty" {
		bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{CallbackQueryID: callback.ID})
		return
	} else if filepath.Ext(callback.Data) != "" {
		// contains extension, handle file instead of folders
		// TODO: make a more robust check, some files don't contain extensions
		// TODO: double-check the callback received to prevent list env files or unauthorized ones
		fileCallback(bot, callback)
		return
	}

	directory := callback.Data
	previousDirectory := getPreviousDirectory(directory)

	buttons := makeButtonsFromFileEntries(directory)
	buttons = prependNavigationButtons(buttons, directory, previousDirectory)
	inlineKeyboard := makeInlineKeyboard(buttons)

	bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{CallbackQueryID: callback.ID})
	bot.EditMessageReplyMarkup(&telego.EditMessageReplyMarkupParams{
		MessageID:   callback.Message.GetMessageID(),
		ChatID:      tutil.ID(callback.Message.GetChat().ID),
		ReplyMarkup: inlineKeyboard,
	})
}

func listCommand(bot *telego.Bot, message telego.Message) {
	directory := rootDirectory
	if admin := isUserAdmin(message); !admin {
		directory = defaultRootDir
	}

	buttons := makeButtonsFromFileEntries(directory)
	buttons = prependNavigationButtons(buttons, directory, directory)
	inlineKeyboard := makeInlineKeyboard(buttons)
	text := "*Click on the buttons below to navigate through folders:*"

	reply := tutil.Message(
		tutil.ID(message.From.ID),
		text,
	).WithReplyMarkup(inlineKeyboard).
		WithParseMode(telego.ModeMarkdownV2)

	_, err := bot.SendMessage(reply)
	if err != nil {
		slog.Error(err.Error())
	}
}

func fileCallback(bot *telego.Bot, callback telego.CallbackQuery) {
	sendID := tutil.ID(callback.From.ID)
	defer bot.AnswerCallbackQuery(&telego.AnswerCallbackQueryParams{CallbackQueryID: callback.ID})

	file, err := os.Open(callback.Data)
	if err != nil {
		slog.Error("Error loading file: " + callback.Data + "|err: " + err.Error())
		bot.SendMessage(tutil.Message(sendID, "Error loading required file"))
	}

	document := tutil.Document(
		sendID,
		tutil.File(file),
	)
	_, err = bot.SendDocument(document)
	if err != nil {
		slog.Error(err.Error())
	}
}

func unknown(bot *telego.Bot, message telego.Message) {
	bot.SendMessage(tutil.Message(
		tutil.ID(message.Chat.ID),
		"Unknown command, use /start",
	))
}
