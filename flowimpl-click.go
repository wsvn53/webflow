package main

import (
	"github.com/chromedp/chromedp"
)

type FlowImplClick FlowImplBase

func (impl *FlowImplClick) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	textSelector := impl.command.Fields[0].ToString()

	err := chromedp.Run(browser.chromeContext,
		chromedp.Click(textSelector, chromedp.ByQuery, chromedp.NodeVisible),
	)
	return err
}

//go:generate make IMPL_TYPE=FlowImplClick gen-impl


func (impl *FlowImplClick) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplClick) Type() FlowImplType {
	return impl.implType
}
