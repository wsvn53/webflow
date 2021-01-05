package main

import (
	"testing"
)

func TestNewBrowser(t *testing.T) {
	flowString := `
	screen  	"1024x800";
	headless 	false;
	timeout 	30000;
	userdata 	"./test/data";
	debug		false;
	log		true;
	open	"https://www.baidu.com/";
	wait	".soutu-btn";
	setvalue  "#kw" "google";
	wait 	1000;
	click 	"#su";
	eval 	"console.log(1234)";
	eval 	$result 	"9 * 123";
	eval 	"console.log($result)";
	print	$result;
	print	$body_text;
	wait	500;
	var 	$atext	" pixel 4a";
	value 	"#kw"	"abc";
	print	$atext;
	keys	"#kw" 	$atext;
	click	"#su";
	wait	15000;
	`
	flow := FlowFromString(flowString)
	b := NewBrowser(flow)
	_ = b.Run()
}
