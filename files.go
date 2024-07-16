package main

import (
	"fmt"
	"log/slog"
	"os"
)

var (
	defaultRootDir = "./testfiles/"
	rootDirectory  string
)

func listFiles(path string) []os.DirEntry {
	entries, err := os.ReadDir(path)
	if err != nil {
		slog.Error(fmt.Sprintf("Error listing path: %s, err: %s", path, err.Error()))
	}
	return entries
}
