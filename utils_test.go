package main

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/mymmrac/telego"
	"github.com/stretchr/testify/assert"
)

func TestIsUserAdmin(t *testing.T) {
	admin := rand.Int63()
	randomID := rand.Int63()
	adminID = strconv.FormatInt(admin, 10)

	type testCases struct {
		name     string
		id       int64
		expected bool
	}
	testcases := []testCases{
		{
			name:     "is admin",
			id:       admin,
			expected: true,
		},
		{
			name:     "is not an admin",
			id:       randomID,
			expected: false,
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			message := &telego.Message{
				From: &telego.User{
					ID: test.id,
				},
			}

			result := isUserAdmin(*message)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestListFiles(t *testing.T) {
	rootDirectory = defaultRootDir

	t.Run("root directory", func(t *testing.T) {
		entries := listFiles(rootDirectory)
		assert.Len(t, entries, 5)
	})

	t.Run("folder with mixed items", func(t *testing.T) {
		entries := listFiles(rootDirectory + "mixed")
		assert.Len(t, entries, 4)

		assert.Equal(t, entries[0].Name(), "666.txt")
		assert.False(t, entries[0].IsDir())

		assert.Equal(t, entries[1].Name(), "itr.jpg")
		assert.False(t, entries[1].IsDir())

		assert.Equal(t, entries[2].Name(), "nested")
		assert.True(t, entries[2].IsDir())

		assert.Equal(t, entries[3].Name(), "nested2")
		assert.True(t, entries[3].IsDir())
	})
}

func TestExtensionToEmoji(t *testing.T) {
	t.Run("return "+TEXT_EMOJI, func(t *testing.T) {
		extensions := []string{".txt", ".csv"}
		for _, e := range extensions {
			t.Run("Extension "+e+" should return "+TEXT_EMOJI, func(t *testing.T) {
				assert.Equal(t, TEXT_EMOJI, extensionToEmoji(e))
			})
		}
	})

	t.Run("return "+PICTURE_EMOJI, func(t *testing.T) {
		extensions := []string{".jpg", ".jpeg", ".png"}
		for _, e := range extensions {
			t.Run("Extension "+e+" should return "+PICTURE_EMOJI, func(t *testing.T) {
				assert.Equal(t, PICTURE_EMOJI, extensionToEmoji(e))
			})
		}
	})

	t.Run("return "+VIDEO_EMOJI, func(t *testing.T) {
		extensions := []string{".mp4", ".mpeg", ".webm"}
		for _, e := range extensions {
			t.Run("Extension "+e+" should return "+VIDEO_EMOJI, func(t *testing.T) {
				assert.Equal(t, VIDEO_EMOJI, extensionToEmoji(e))
			})
		}
	})

	t.Run("return "+AUDIO_EMOJI, func(t *testing.T) {
		extensions := []string{".ogg", ".mp3", ".aac"}
		for _, e := range extensions {
			t.Run("Extension "+e+" should return "+AUDIO_EMOJI, func(t *testing.T) {
				assert.Equal(t, AUDIO_EMOJI, extensionToEmoji(e))
			})
		}
	})

	t.Run("return "+CODE_EMOJI, func(t *testing.T) {
		extensions := []string{".html", ".css", ".js"}
		for _, e := range extensions {
			t.Run("Extension "+e+" should return "+CODE_EMOJI, func(t *testing.T) {
				assert.Equal(t, CODE_EMOJI, extensionToEmoji(e))
			})
		}
	})
}
