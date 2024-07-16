package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
