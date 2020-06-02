package main

import (
	"fmt"
	"os"
)

func assertErr(tag string, err error) {
	if err == nil {
		return
	}

	fmt.Printf("[!] %s: %s\n", tag, err)
	os.Exit(1)
}
