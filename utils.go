package main

import (
	"log/slog"
	"strconv"

	"github.com/mymmrac/telego"

	tutil "github.com/mymmrac/telego/telegoutil"
)

func isUserAdmin(update telego.Update) bool {
	return strconv.FormatInt(update.Message.From.ID, 10) == adminID
}

func checkForAdminStatus(update telego.Update) bool {
	if !isUserAdmin(update) {
		msg := tutil.Message(
			tutil.ID(update.Message.Chat.ID),
			"User unauthorized",
		)
		_, err := bot.SendMessage(msg)
		if err != nil {
			slog.Error(err.Error())
		}
		return false
	}
	slog.Info("Authorized request")
	return true
}
