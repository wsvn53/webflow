package main

import (
	"github.com/chromedp/chromedp"
	"reflect"
	"strings"
)

type FlowImplKeys FlowImplBase

func (impl FlowImplKeys) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplKeys) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	textSelector := impl.command.Fields[0].ToString()
	typingKeys := impl.command.Fields[1].ToString()

	if strings.HasPrefix(typingKeys, "$") {
		textValues, ok := browser.variableMaps[typingKeys]
		if ok {
			typingKeys = textValues
		}
	}

	err := chromedp.Run(browser.chromeContext,
		chromedp.SendKeys(textSelector, typingKeys, chromedp.ByQuery),
	)
	return err
}

//go:generate make IMPL_TYPE=FlowImplKeys gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplKeys{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplKeys) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplKeys) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplKeys) Command() *FlowCommand {
	return impl.command
}
