package main

import (
	"github.com/chromedp/chromedp"
	"strings"
)

type FlowImplEval FlowImplBase

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


func (impl *FlowImplEval) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplEval) Type() FlowImplType {
	return impl.implType
}
