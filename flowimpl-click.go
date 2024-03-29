package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplClick FlowImplBase

func (impl *FlowImplClick) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplClick) Usage() string {
	return impl.Name() + " <selector>		Click element by <selector>, like document.querySelector"
}

func (impl *FlowImplClick) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	textSelector := impl.command.GetFieldString(0)

	err := chromedp.Run(browser.chromeContext,
		chromedp.Click(textSelector, chromedp.ByQuery, chromedp.NodeVisible, chromedp.FromNode(browser.switchNode)),
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

func (impl *FlowImplClick) Clone() IFlowImpl {
	c := &FlowImplClick{}
	_ = copier.Copy(c, impl)
	return c
}
