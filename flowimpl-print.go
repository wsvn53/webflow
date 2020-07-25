package main

import (
	"fmt"
	"reflect"
	"strings"
)

type FlowImplPrint FlowImplBase

func (impl *FlowImplPrint) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplPrint) Do(args...interface{}) error {
	browser := args[0].(*Browser)

	out := impl.command.Fields[0].ToString()
	if _, ok := browser.variableMaps[out]; ok && strings.HasPrefix(out, "$") {
		fmt.Println(browser.variableMaps[out])
		return nil
	}
	fmt.Println(out)

	return nil
}

//go:generate make IMPL_TYPE=FlowImplPrint gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplPrint{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplPrint) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplPrint) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplPrint) Command() *FlowCommand {
	return impl.command
}
