package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplHeadless FlowImplBase

func (impl *FlowImplHeadless) Type() FlowImplType {
	return FlowImplTypeFlag
}

func (impl *FlowImplHeadless) Usage() string {
	return impl.Name() + " <true|false> 	Set headless flag to chromedp, default is true means don't show browser window"
}

func (impl *FlowImplHeadless) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}

	opt := args[0].(*chromedp.ExecAllocatorOption)
	if len(impl.command.Fields) == 0 {
		*opt = chromedp.Flag("headless", true)
	}

	switch impl.command.Fields[0].ToString() {
	case "true":
		*opt = chromedp.Flag("headless", true)
	case "false":
		*opt = chromedp.Flag("headless", false)
	}

	return nil
}

//go:generate make IMPL_TYPE=FlowImplHeadless gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplHeadless{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplHeadless) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplHeadless) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplHeadless) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplHeadless) Clone() IFlowImpl {
	c := &FlowImplHeadless{}
	_ = copier.Copy(c, impl)
	return c
}
