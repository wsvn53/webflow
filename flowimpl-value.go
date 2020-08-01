package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplValue FlowImplBase

func (impl FlowImplValue) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplValue) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	textSelector := impl.command.Fields[0].ToString()
	toValue := impl.command.Fields[1].ToString()

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

//go:generate make IMPL_TYPE=FlowImplValue gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplValue{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplValue) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplValue) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplValue) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplValue) Clone() IFlowImpl {
	c := &FlowImplValue{}
	_ = copier.Copy(c, impl)
	return c
}
