package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func assertErr(tag string, err error) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, "[!] %s: %s\n", tag, err)
	os.Exit(1)
}

func resolvePath(path string) string {
	switch {
	case strings.HasPrefix(path, "~") && len(path) > 2:
		homeDir, _ := os.UserHomeDir()
		if homeDir == "" {
			homeDir = "~"
		}
		path = filepath.Join(homeDir, path[2:])
	case strings.HasPrefix(path, ".") && len(path) > 2:
		workDir, _ := os.Getwd()
		path = filepath.Join(workDir, path[2:])
	}
	return path
}