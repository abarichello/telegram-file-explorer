package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mymmrac/telego"

	_ "github.com/joho/godotenv/autoload"
	thandler "github.com/mymmrac/telego/telegohandler"
)

var (
	logger     = slog.New(tint.NewHandler(os.Stderr, nil))
	bot        *telego.Bot
	botHandler *thandler.BotHandler
)

func setup() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	loadEnvs()

	b, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		slog.Error(err.Error())
	}
	bot = b
	slog.Info("Started bot")
}

func main() {
	setup()
	slog.Info("Finished setup")

	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		slog.Error(err.Error())
	}
	defer bot.StopLongPolling()

	bh, err := thandler.NewBotHandler(bot, updates)
	if err != nil {
		slog.Error(err.Error())
	}
	botHandler = bh
	defer botHandler.Stop()

	setMyCommands()
	registerHandlers()
	botHandler.Start()
}
