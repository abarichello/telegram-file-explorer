package main

import (
	"fmt"
	"log/slog"
	"mime"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/mymmrac/telego"

	tutil "github.com/mymmrac/telego/telegoutil"
)

var (
	adminID        string
	botToken       string
	rootDirectory  string
	defaultRootDir = "testfiles"
)

func loadEnvs() {
	botToken = os.Getenv("BOT_TOKEN")
	if botToken == "" {
		slog.Error("No BOT_TOKEN env exported")
	}

	adminID = os.Getenv("ADMIN_ID")
	if adminID == "" {
		slog.Error("No ADMIN_ID environment found")
	}

	envRoot := os.Getenv("DIRECTORY_ROOT")
	if envRoot == "" {
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
	// TODO: sort directories first
	entries, err := os.ReadDir(path)
	if err != nil {
		slog.Error(fmt.Sprintf("Error listing path: %s, err: %s", path, err.Error()))
	}
	return entries
}

func getPreviousDirectory(path string) string {
	if path == strings.TrimRight(rootDirectory, "/") {
		return rootDirectory
	}

	normalizedPath := strings.TrimRight(path, "/")
	previous := filepath.Dir(normalizedPath)
	return previous
}

func extensionToEmoji(ext string) string {
	mimeMap := map[string]string{
		"text/plain": TEXT_EMOJI,
		"text/csv":   TEXT_EMOJI,

		"image/gif":     PICTURE_EMOJI,
		"image/jpeg":    PICTURE_EMOJI,
		"image/png":     PICTURE_EMOJI,
		"image/svg+xml": PICTURE_EMOJI,

		"video/mpeg": VIDEO_EMOJI,
		"video/mp4":  VIDEO_EMOJI,
		"video/webm": VIDEO_EMOJI,

		"audio/aac":       AUDIO_EMOJI,
		"audio/wav":       AUDIO_EMOJI,
		"audio/wave":      AUDIO_EMOJI,
		"audio/x-wav":     AUDIO_EMOJI,
		"audio/x-pn-wav":  AUDIO_EMOJI,
		"audio/ogg":       AUDIO_EMOJI,
		"audio/mpeg":      AUDIO_EMOJI,
		"application/ogg": AUDIO_EMOJI,

		"text/css":        CODE_EMOJI,
		"text/javascript": CODE_EMOJI,
		"text/htm":        CODE_EMOJI,
		"text/html":       CODE_EMOJI,
	}

	for k, v := range mimeMap {
		mimeExtensions, err := mime.ExtensionsByType(k)
		if err != nil {
			slog.Error(err.Error())
		}
		if slices.Contains(mimeExtensions, ext) {
			return v
		}
	}

	return GENERIC_FILE_EMOJI
}
