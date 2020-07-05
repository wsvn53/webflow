package main

import (
	"fmt"
	"os"
)

func assertErr(tag string, err error) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, "[!] %s: %s\n", tag, err)
	os.Exit(1)
}
