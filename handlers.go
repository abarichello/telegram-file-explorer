package main

import (
	"fmt"
	"log/slog"

	"github.com/mymmrac/telego"

	thandler "github.com/mymmrac/telego/telegohandler"
	tutil "github.com/mymmrac/telego/telegoutil"
)

func setMyCommands() {
	bot.SetMyCommands(&telego.SetMyCommandsParams{
		Commands: []telego.BotCommand{
			{Command: "start", Description: "Starts the bot"},
			{Command: "list", Description: "Lists root directory"},
			{Command: "help", Description: "Shows how to use"},
		},
	})
}

func registerHandlers() {
	botHandler.HandleMessage(start, thandler.CommandEqual("start"))
	botHandler.HandleMessage(list, thandler.CommandEqual("list"))
	botHandler.HandleMessage(unknown)
}

func start(bot *telego.Bot, message telego.Message) {
	if !checkForAdminStatus(message) {
		return
	}

	msg := tutil.Message(
		tutil.ID(message.Chat.ID),
		fmt.Sprintf("Oi %d", message.From.ID),
	)
	_, err := bot.SendMessage(msg)
	if err != nil {
		slog.Error(err.Error())
	}
}

func list(bot *telego.Bot, message telego.Message) {
	directory := rootDirectory
	if admin := isUserAdmin(message); !admin {
		directory = defaultRootDir
	}

	var buttons []telego.InlineKeyboardButton
	entries := listFiles(directory)
	for _, entry := range entries {
		fn := fileButton
		suffix := ""
		if entry.IsDir() {
			fn = folderButton
			suffix = "/"
		}
		newPath := directory + entry.Name() + suffix
		buttons = append(buttons, fn(entry.Name(), newPath))
	}

	inlineKeyboard := makeInlineKeyboard(buttons)
	text := "*Listing directory:* " + directory + "\n"

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

func unknown(bot *telego.Bot, message telego.Message) {
	bot.SendMessage(tutil.Message(
		tutil.ID(message.Chat.ID),
		"Unknown command, use /start",
	))
}
