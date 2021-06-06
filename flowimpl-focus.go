package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplFocus FlowImplBase

func (impl *FlowImplFocus) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplFocus) Usage() string {
	return impl.Name() + " <selector>		Focus element by <selector>, like document.querySelector"
}

func (impl *FlowImplFocus) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	textSelector := impl.command.GetFieldString(0)

	err := chromedp.Run(browser.chromeContext,
		chromedp.Focus(textSelector, chromedp.ByQuery, chromedp.NodeVisible, chromedp.FromNode(browser.switchNode)),
	)
	return err
}

//go:generate make IMPL_TYPE=FlowImplFocus gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplFocus{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplFocus) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplFocus) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplFocus) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplFocus) Clone() IFlowImpl {
	c := &FlowImplFocus{}
	_ = copier.Copy(c, impl)
	return c
}
