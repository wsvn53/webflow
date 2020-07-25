package main

import (
	"errors"
	"fmt"
	"github.com/chromedp/chromedp"
	"reflect"
	"strings"
)

type FlowImplValue FlowImplBase

func (impl *FlowImplValue) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplValue) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	variableName := impl.command.Fields[0].ToString()
	textSelector := impl.command.Fields[1].ToString()

	if strings.HasPrefix(variableName, "$") == false {
		return errors.New(fmt.Sprintf("%s is not a valid variable.", variableName))
	}

	var textContent string
	err := chromedp.Run(browser.chromeContext,
		chromedp.TextContent(textSelector, &textContent, chromedp.ByQuery),
	)
	browser.SetVariable(variableName, textContent)
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
