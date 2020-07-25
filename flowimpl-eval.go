package main

import (
	"github.com/chromedp/chromedp"
	reflect "reflect"
	"strings"
)

type FlowImplEval FlowImplBase

func (impl *FlowImplEval) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplEval) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	script := impl.command.Fields[0].ToString()
	var variableName string
	if len(impl.command.Fields) > 1 {
		variableName = impl.command.Fields[0].ToString()
		script = impl.command.Fields[1].ToString()
	}

	var result interface{}
	_ = chromedp.Run(browser.chromeContext,
		chromedp.Evaluate(script, &result),
	)

	if strings.HasPrefix(variableName, "$") {
		browser.SetVariable(variableName, result)
	}

	return nil
}

//go:generate make IMPL_TYPE=FlowImplEval gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplEval{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplEval) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplEval) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplEval) Command() *FlowCommand {
	return impl.command
}
