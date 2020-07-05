package main

import (
	"testing"
)

func TestNewBrowser(t *testing.T) {
	flowString := `
	screen  	"1024x800";
	headless 	false;
	timeout 	30000;
	debug		false;
	log		false;
	open	"https://www.baidu.com/";
	wait	".soutu-btn";
	text 	$body_text  "#s-top-left";
	eval 	"console.log($body_text);";
	eval 	$result 	"9 * 123";
	eval 	"console.log($result)";
	print	$result;
	print	$body_text;
	wait	500;
	text 	$atext	"a";
	print	$atext;
	click	".soutu-btn";
	wait	2000;
	`
	flow := FlowFromString(flowString)
	b := NewBrowser(flow)
	_ = b.Run()
}
