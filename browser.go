package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
)

type Browser struct {
	webFlow 		*Flow
	chromeContext 	context.Context
	chromeCancel 	context.CancelFunc
}

func NewBrowser(flow *Flow) *Browser {
	chromeOpts := append([]chromedp.ExecAllocatorOption{}, chromedp.DefaultExecAllocatorOptions[:]...)
	flow.WalkByType(FlowImplTypeFlag, func(i int, impl IFlowImpl, stop *bool) {
		var opt chromedp.ExecAllocatorOption
		err := impl.Do(&opt)
		if opt != nil {
			chromeOpts = append(chromeOpts, opt)
		}
		assertErr("Setup", err)
	})

	c := context.Background()
	ctx, cancel := chromedp.NewExecAllocator(c, chromeOpts...)
	ctx, cancel = chromedp.NewContext(ctx,
		chromedp.WithLogf(log.Printf),
		chromedp.WithDebugf(log.Printf),
		chromedp.WithErrorf(log.Printf))

	return &Browser{
		webFlow: flow,
		chromeContext: ctx,
		chromeCancel: cancel,
	}
}

func (browser *Browser) Run() error {
	browser.webFlow.Walk(func(i int, impl IFlowImpl, stop *bool) {
		if impl.Type() == FlowImplTypeFlag {
			return
		}
		fmt.Println("==> Running Task:", *impl.Command().Name)
		err := impl.Do(browser)
		assertErr("Run", err)
	})

	return nil
}
