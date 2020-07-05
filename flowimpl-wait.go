package main

import (
	"github.com/chromedp/chromedp"
	"strconv"
	"time"
)

type FlowImplWait FlowImplBase

func (impl *FlowImplWait) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	waitTarget := impl.command.Fields[0].ToString()
	timeout, err := strconv.Atoi(waitTarget)
	if err == nil {
		// wait for timeout
		err = chromedp.Run(browser.chromeContext,
			chromedp.Sleep(time.Millisecond * time.Duration(timeout)))
		return err
	}

	err = chromedp.Run(browser.chromeContext,
		chromedp.WaitReady(waitTarget, chromedp.ByQuery),
	)
	return err
}

//go:generate make IMPL_TYPE=FlowImplWait gen-impl


func (impl *FlowImplWait) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplWait) Type() FlowImplType {
	return impl.implType
}
