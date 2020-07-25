package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type FlowImplLog FlowImplBase

func (impl *FlowImplLog) Type() FlowImplType {
	return FlowImplTypeOP
}

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

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplLog{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplLog) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplLog) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplLog) Command() *FlowCommand {
	return impl.command
}
