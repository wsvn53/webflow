package main

import (
	"fmt"
	"testing"
)

func TestNewBrowser(t *testing.T) {
	flowString := `
	screen  "1024x800";
	headless false;
	timeout 	10;
	open	"https://www.example.com/";
	text 	$body_text  "body div";
	eval 	"console.log(\"hello world!\" + body_text);";
	eval 	$result 	"(1 + 1).toString()";
	eval 	"console.log(result)";
	`
	flow := FlowFromString(flowString)
	b := NewBrowser(flow)
	fmt.Println(b)

	_ = b.Run()
}
