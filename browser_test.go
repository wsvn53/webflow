package main

import (
	"testing"
)

func TestNewBrowser(t *testing.T) {
	flowString := `
	screen  "1024x800";
	headless false;
	timeout 	10;
	debug		false;
	log		false;
	open	"https://www.example.com/";
	text 	$body_text  "body div";
	eval 	"console.log(\"hello world!\" + $body_text);";
	eval 	$result 	"9 * 123";
	eval 	"console.log($result)";
	print	$result;
	print	$body_text;
	`
	flow := FlowFromString(flowString)
	b := NewBrowser(flow)
	_ = b.Run()
}
