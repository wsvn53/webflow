package main

import (
	"fmt"
	"testing"
)

func TestNewFlowImpl(t *testing.T) {
	flowString := `open "https://www.example.com/"`
	flow := FlowFromString(flowString)
	cmd := flow.CommandByName("open")

	impl := NewFlowImpl(cmd)
	if impl == nil {
		t.Error("NewFlowImpl not as expected.")
		return
	}

	cmd = impl.Command()
	if cmd == nil {
		t.Error("Get command's impls failed.")
		return
	}
	fmt.Println("==>", *cmd.Name, cmd.FieldsString())
}
