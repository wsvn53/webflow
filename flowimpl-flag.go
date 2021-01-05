package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplFlag FlowImplBase

func (impl *FlowImplFlag) Type() FlowImplType {
	return FlowImplTypeFlag
}

func (impl *FlowImplFlag) Usage() string {
	return impl.Name() + " <string> 	Custom browser flags."
}

func (impl *FlowImplFlag) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	opt := args[0].(*chromedp.ExecAllocatorOption)
	flagName := impl.command.GetFieldString(0)
	flagValue := impl.command.GetFieldString(1)
	switch flagValue {
	case "true":
		*opt = chromedp.Flag(flagName, true)
	case "false":
		*opt = chromedp.Flag(flagName, false)
	default:
		*opt = chromedp.Flag(flagName, flagValue)
	}
	return nil
}

//go:generate make IMPL_TYPE=FlowImplFlag gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplFlag{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplFlag) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplFlag) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplFlag) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplFlag) Clone() IFlowImpl {
	c := &FlowImplFlag{}
	_ = copier.Copy(c, impl)
	return c
}
