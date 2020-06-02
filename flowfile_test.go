package main

import (
	"fmt"
	"testing"
)

func TestFlowFromString(t *testing.T) {
	flowString := `
open	"https://www.baidu.com/";
timeout 	30;
text $body_text  "body.div";
eval "console.log('hello world!')";
eval $result 	"1 + 1";
`
	flow := FlowFromString(flowString)

	for _, c := range flow.FlowCommands {
		fmt.Println(*c.Name, "=>", len(c.Fields))
		for _, f := range c.Fields {
			fmt.Println(" - ", f.ToString())
		}
	}
}

