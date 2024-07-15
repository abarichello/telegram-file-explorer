package main

import (
	"fmt"
	"log/slog"

	"github.com/mymmrac/telego"

	thandler "github.com/mymmrac/telego/telegohandler"
	tutil "github.com/mymmrac/telego/telegoutil"
)

func registerHandlers(botHandler *thandler.BotHandler) {
	botHandler.Handle(start)
}

func start(bot *telego.Bot, update telego.Update) {
	if !checkForAdminStatus(update) {
		return
	}

	msg := tutil.Message(
		tutil.ID(update.Message.Chat.ID),
		fmt.Sprintf("Oi %d", update.Message.From.ID),
	)
	_, err := bot.SendMessage(msg)
	if err != nil {
		slog.Error(err.Error())
	}
}
