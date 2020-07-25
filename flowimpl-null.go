package main

import (
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplNull FlowImplBase

func (impl *FlowImplNull) Type() FlowImplType {
	return FlowImplTypeNull
}

func (impl *FlowImplNull) Do(args...interface{}) error {
	// null task do nothing
	return nil
}

//go:generate make IMPL_TYPE=FlowImplNull gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplNull{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplNull) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplNull) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplNull) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplNull) Clone() IFlowImpl {
	c := &FlowImplNull{}
	_ = copier.Copy(c, impl)
	return c
}
