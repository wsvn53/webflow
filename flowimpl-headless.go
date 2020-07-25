package main

import (
	"github.com/chromedp/chromedp"
	"reflect"
	"strings"
)

type FlowImplHeadless FlowImplBase

func (impl *FlowImplHeadless) Type() FlowImplType {
	return FlowImplTypeFlag
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
		fallthrough
	case "YES":
		fallthrough
	case "1":
		*opt = chromedp.Flag("headless", true)
	case "false":
		fallthrough
	case "NO":
		fallthrough
	case "0":
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
