package main

import (
	"github.com/chromedp/chromedp"
	"reflect"
	"strings"
)

type FlowImplOpen FlowImplBase

func (impl *FlowImplOpen) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplOpen) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	url := impl.command.Fields[0].ToString()
	err := chromedp.Run(browser.chromeContext, chromedp.Navigate(url))
	return err
}

//go:generate make IMPL_TYPE=FlowImplOpen gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplOpen{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplOpen) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplOpen) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplOpen) Command() *FlowCommand {
	return impl.command
}
