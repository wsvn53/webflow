package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplSetValue FlowImplBase

func (impl *FlowImplSetValue) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplSetValue) Usage() string {
	return impl.Name() + " <selector> <string|$variable> 	Set the value of element <selector> with <string> or $variable"
}

func (impl *FlowImplSetValue) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	textSelector := impl.command.GetFieldString(0)
	toValue := impl.command.GetFieldString(1)

	if strings.HasPrefix(toValue, "$") {
		textValues, ok := browser.variableMaps[toValue]
		if ok {
			toValue = textValues
		}
	}

	err := chromedp.Run(browser.chromeContext,
		chromedp.SetValue(textSelector, toValue, chromedp.ByQuery),
	)
	return err
}

//go:generate make IMPL_TYPE=FlowImplSetValue gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplSetValue{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplSetValue) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplSetValue) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplSetValue) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplSetValue) Clone() IFlowImpl {
	c := &FlowImplSetValue{}
	_ = copier.Copy(c, impl)
	return c
}
