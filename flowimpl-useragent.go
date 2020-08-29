package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplUserAgent FlowImplBase

func (impl *FlowImplUserAgent) Type() FlowImplType {
	return FlowImplTypeFlag
}

func (impl *FlowImplUserAgent) Usage() string {
	return impl.Name() + " <string> 	Setup browser UserAgent string"
}

func (impl *FlowImplUserAgent) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	opt := args[0].(*chromedp.ExecAllocatorOption)
	*opt = chromedp.UserAgent(impl.command.GetFieldString(0))
	return nil
}

//go:generate make IMPL_TYPE=FlowImplUserAgent gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplUserAgent{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplUserAgent) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplUserAgent) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplUserAgent) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplUserAgent) Clone() IFlowImpl {
	c := &FlowImplUserAgent{}
	_ = copier.Copy(c, impl)
	return c
}
