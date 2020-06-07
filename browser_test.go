package main

import (
	"fmt"
	"testing"
)

func TestNewBrowser(t *testing.T) {
	flowString := `
	screen  "1024x800";
	headless false;
	open	"https://www.example.com/";
	timeout 	30;
	text $body_text  "body.div";
	eval "console.log('hello world!');";
	eval $result 	"1 + 1";
	`
	flow := FlowFromString(flowString)
	b := NewBrowser(flow)
	fmt.Println(b)

	_ = b.Run()
}
