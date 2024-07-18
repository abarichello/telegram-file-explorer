package main

import (
	"math"
	"path/filepath"

	"github.com/enescakir/emoji"
	"github.com/mymmrac/telego"

	tutil "github.com/mymmrac/telego/telegoutil"
)

// maxColumns is the max amount of columns of buttons to be showed with the inline keyboard
const maxColumns = 4

var (
	BACK_EMOJI      = emoji.LeftArrow.String()
	FOLDER_EMOJI    = emoji.FileFolder.String()
	DIRECTORY_EMOJI = emoji.FileCabinet.String()

	TEXT_EMOJI         = emoji.PageFacingUp.String()
	PICTURE_EMOJI      = emoji.FramedPicture.String()
	VIDEO_EMOJI        = emoji.FilmFrames.String()
	AUDIO_EMOJI        = emoji.MusicalNotes.String()
	CODE_EMOJI         = emoji.Keyboard.String()
	GENERIC_FILE_EMOJI = emoji.PageWithCurl.String()
)

func backButton(path string) telego.InlineKeyboardButton {
	return tutil.InlineKeyboardButton(BACK_EMOJI + " Go back").WithCallbackData(path)
}

func folderButton(name, path string) telego.InlineKeyboardButton {
	return tutil.InlineKeyboardButton(FOLDER_EMOJI + " " + name).WithCallbackData(path)
}

func currentDirButton(name string) telego.InlineKeyboardButton {
	return tutil.InlineKeyboardButton("Directory: " + name).WithCallbackData("empty")
}

func fileButton(name, path string) telego.InlineKeyboardButton {
	extension := filepath.Ext(path)
	emoji := extensionToEmoji(extension)
	return tutil.InlineKeyboardButton(emoji + " " + name).WithCallbackData(path)
}

func makeButtonsFromFileEntries(path string) []telego.InlineKeyboardButton {
	var buttons []telego.InlineKeyboardButton
	entries := listFiles(path)
	for _, entry := range entries {
		fn := fileButton
		suffix := ""
		if entry.IsDir() {
			fn = folderButton
			suffix = "/"
		}
		newPath := path + "/" + entry.Name() + suffix
		buttons = append(buttons, fn(entry.Name(), newPath))
	}
	return buttons
}

func prependNavigationButtons(buttons []telego.InlineKeyboardButton, directory, previousDir string) []telego.InlineKeyboardButton {
	previousButton := backButton(previousDir)
	dirButton := currentDirButton(directory)
	buttons = append([]telego.InlineKeyboardButton{previousButton, dirButton}, buttons...)
	return buttons
}

// Returns an inline keyboard based on an array of buttons
// Organizes them neatly according to maxColumns
func makeInlineKeyboard(buttons []telego.InlineKeyboardButton) *telego.InlineKeyboardMarkup {
	rows := [][]telego.InlineKeyboardButton{
		{buttons[0], buttons[1]},
	}
	i := 2 // skip first two buttons
	for ; i < len(buttons); i += maxColumns {
		endSlice := int(math.Min(float64(i+maxColumns), float64(len(buttons)))) // cap slice index
		rows = append(rows, buttons[i:endSlice])
	}
	inlineKeyboard := tutil.InlineKeyboard(rows...)
	return inlineKeyboard
}
