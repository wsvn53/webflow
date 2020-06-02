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
	}
	fmt.Println("==>", *impl.Command().Name, impl.Command().FieldsString())
}
