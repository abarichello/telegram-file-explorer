package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"

	tutil "github.com/mymmrac/telego/telegoutil"
)

var (
	adminID        string
	rootDirectory  string
	defaultRootDir = "./testfiles/"
)

func loadEnvs() {
	envMap, err := godotenv.Read()
	if err != nil {
		log.Fatal("Env loading error: ", err)
	}

	envID, ok := envMap["ADMIN_ID"]
	if !ok {
		slog.Error("No ADMIN_ID environment found")
	}
	adminID = envID

	envRoot, ok := envMap["DIRECTORY_ROOT"]
	if !ok {
		slog.Info("No DIRECTORY_ROOT environment found, using test dir as default")
		rootDirectory = defaultRootDir
	} else {
		rootDirectory = envRoot
	}
}

func isUserAdmin(message telego.Message) bool {
	return strconv.FormatInt(message.From.ID, 10) == adminID
}

func checkForAdminStatus(message telego.Message) bool {
	if !isUserAdmin(message) {
		slog.Info("Unauthorized request")
		msg := tutil.Message(
			tutil.ID(message.Chat.ID),
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

func listFiles(path string) []os.DirEntry {
	entries, err := os.ReadDir(path)
	if err != nil {
		slog.Error(fmt.Sprintf("Error listing path: %s, err: %s", path, err.Error()))
	}
	return entries
}
