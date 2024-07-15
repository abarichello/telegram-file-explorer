package main

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
	"github.com/mymmrac/telego"

	thandler "github.com/mymmrac/telego/telegohandler"
)

var (
	logger = slog.New(tint.NewHandler(os.Stderr, nil))
	bot    *telego.Bot

	adminID string
)

func setup() {
	envMap, err := godotenv.Read()
	if err != nil {
		log.Fatal("Env loading error: ", err)
	}

	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))
	slog.Info("Started bot")

	b, err := telego.NewBot(envMap["BOT_TOKEN"], telego.WithDefaultDebugLogger())
	if err != nil {
		slog.Error(err.Error())
	}
	bot = b

	envID, ok := envMap["ADMIN_ID"]
	if !ok {
		slog.Error("No ADMIN_ID environment found")
	}
	adminID = envID
}

func main() {
	setup()
	slog.Info("Finished setup")

	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		slog.Error(err.Error())
	}
	defer bot.StopLongPolling()

	botHandler, err := thandler.NewBotHandler(bot, updates)
	if err != nil {
		slog.Error(err.Error())
	}
	defer botHandler.Stop()

	registerHandlers(botHandler)
	botHandler.Start()
}
