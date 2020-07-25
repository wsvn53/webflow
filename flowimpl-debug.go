package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"log"
	"reflect"
	"strings"
)

type FlowImplDebug FlowImplBase

func (impl *FlowImplDebug) Type() FlowImplType {
	return FlowImplTypeLog
}

func (impl *FlowImplDebug) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	opt := args[0].(*chromedp.ContextOption)

	if impl.command.Fields[0].ToString() == "true" {
		*opt = chromedp.WithDebugf(log.Printf)
	}

	return nil
}

//go:generate make IMPL_TYPE=FlowImplDebug gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplDebug{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplDebug) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplDebug) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplDebug) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplDebug) Clone() IFlowImpl {
	c := &FlowImplDebug{}
	_ = copier.Copy(c, impl)
	return c
}
