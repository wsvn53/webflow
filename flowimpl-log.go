package main

import (
	"fmt"
	"os"
)

type FlowImplLog FlowImplBase

func (impl *FlowImplLog) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}

	browser := args[0].(*Browser)
	logSwitch := impl.command.Fields[0].ToString()
	if logSwitch == "true" {
		var logFunc = func(a ...interface{}) (n int, err error) {
			return fmt.Fprintln(os.Stderr, a...)
		}
		browser.logFunc = &logFunc
	}

	return nil
}

//go:generate make IMPL_TYPE=FlowImplLog gen-impl


func (impl *FlowImplLog) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplLog) Type() FlowImplType {
	return impl.implType
}
