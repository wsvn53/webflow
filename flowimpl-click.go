package main

import (
	"github.com/chromedp/chromedp"
	"reflect"
	"strings"
)

type FlowImplClick FlowImplBase

func (impl *FlowImplClick) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplClick) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	textSelector := impl.command.Fields[0].ToString()

	err := chromedp.Run(browser.chromeContext,
		chromedp.Click(textSelector, chromedp.ByQuery, chromedp.NodeVisible),
	)
	return err
}

//go:generate make IMPL_TYPE=FlowImplClick gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplClick{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplClick) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplClick) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplClick) Command() *FlowCommand {
	return impl.command
}
