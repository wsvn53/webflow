package main

import (
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	reflect "reflect"
	"strings"
)

type FlowImplEval FlowImplBase

func (impl *FlowImplEval) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplEval) Usage() string {
	return impl.Name() + " [$variable] <{Javascript Code}> 	Evaluate Javascript code and save result to $variable"
}

func (impl *FlowImplEval) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	script := impl.command.GetFieldString(0)
	var variableName string
	if len(impl.command.Fields) > 1 {
		variableName = impl.command.GetFieldString(0)
		script = impl.command.GetFieldString(1)
	}

	var result interface{}
	_ = chromedp.Run(browser.chromeContext,
		chromedp.Evaluate(script, &result),
	)

	if strings.HasPrefix(variableName, "$") {
		browser.SetVariable(variableName, result)
	}

	if variableName == "" {
		fmt.Println(result)
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

func (impl *FlowImplEval) Clone() IFlowImpl {
	c := &FlowImplEval{}
	_ = copier.Copy(c, impl)
	return c
}
