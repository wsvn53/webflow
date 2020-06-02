package main

import (
	"fmt"
	"testing"
)

func TestFlowFromString(t *testing.T) {
	flowString := `
	open	"https://www.example.com/";
	timeout 	30;
	text $body_text  "body.div";
	eval "console.log('hello world!');";
	eval $result 	"1 + 1";
	`
	flow := FlowFromString(flowString)

	// test full flow parsing
	cmd := flow.CommandByName("open")
	fmt.Println("==>", *cmd.Name, cmd.Fields[0].ToString())
	if cmd == nil || cmd.Fields[0].ToString() != "https://www.example.com/" {
		t.Error("Flow parse not as expected!")
	}

	// test walk commands
	flow.WalkCommands(func(i int, cmd *FlowCommand, stop *bool) {
		fmt.Printf("> Step %d: %s args( %s )\n", i + 1, *cmd.Name, cmd.FieldsString())
	})

	// test raw string parsing
	rawString := "console.log('rawstring')\nconsole.log('line 2')"
	flowRawstring := "eval `" + rawString + "`"
	flow = FlowFromString(flowRawstring)
	evalCmd := flow.CommandByName("eval")
	if evalCmd == nil || evalCmd.Fields[0].ToString() != rawString {
		t.Error("RawString parse not as expected!")
	}
	fmt.Printf("==> RawString:\n%s\n", evalCmd.Fields[0].ToString())
}

