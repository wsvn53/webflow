package main

import (
	"errors"
	"fmt"
	"github.com/chromedp/chromedp"
	"strings"
)

/* Impl for all DOM operation commands */

// impl [text] command
type FlowImplText FlowImplBase

func (impl *FlowImplText) Do(args...interface{}) error {
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

func (impl *FlowImplText) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplText) Type() FlowImplType {
	return impl.implType
}
