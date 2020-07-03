package main

import "github.com/chromedp/chromedp"

/* Impl for browser navigation operation commands */

// impl [open] command
type FlowImplOpen FlowImplBase

func (impl *FlowImplOpen) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	url := impl.command.Fields[0].ToString()
	err := chromedp.Run(browser.chromeContext, chromedp.Navigate(url))
	return err
}

func (impl *FlowImplOpen) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplOpen) Type() FlowImplType {
	return impl.implType
}
