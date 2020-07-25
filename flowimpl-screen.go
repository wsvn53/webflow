package main

import (
	"reflect"
	"strings"
)

type FlowImplScreen FlowImplBase

func (impl *FlowImplScreen) Type() FlowImplType {
	return FlowImplTypeFlag
}

func (impl *FlowImplScreen) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	return nil
}

//go:generate make IMPL_TYPE=FlowImplScreen gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplScreen{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplScreen) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplScreen) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplScreen) Command() *FlowCommand {
	return impl.command
}
