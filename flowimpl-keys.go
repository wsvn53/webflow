package main

import (
	"github.com/chromedp/chromedp"
	"strings"
)

type FlowImplKeys FlowImplBase

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

//go:generate make IMPL_TYPE=FlowImplText gen-impl


func (impl *FlowImplKeys) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplKeys) Type() FlowImplType {
	return impl.implType
}
