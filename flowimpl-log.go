package main

import (
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplLog FlowImplBase

func (impl *FlowImplLog) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplLog) Usage() string {
	return impl.Name() + " <true|false> 	Set enable/disable flow verbose log."
}

func (impl *FlowImplLog) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}

	browser := args[0].(*Browser)
	logSwitch := impl.command.GetFieldString(0)
	if logSwitch == "true" {
		browser.setLogEnable(true)
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

func (impl *FlowImplLog) Clone() IFlowImpl {
	c := &FlowImplLog{}
	_ = copier.Copy(c, impl)
	return c
}
