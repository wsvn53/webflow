package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
)

type Browser struct {
	webFlow 		*Flow
	chromeContext 	context.Context
	chromeCancel 	context.CancelFunc

	variableMaps	map[string]string
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
		//chromedp.WithDebugf(log.Printf),
		chromedp.WithErrorf(log.Printf))

	return &Browser{
		webFlow: flow,
		chromeContext: ctx,
		chromeCancel: cancel,
		variableMaps: map[string]string{},
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

func (browser *Browser) SetVariable(name, value string) {
	browser.variableMaps[name] = value

	// store variable to web context
	escapedValue := strings.ReplaceAll(value, "\n", "\\n")
	escapedValue = strings.ReplaceAll(escapedValue, "\r", "\\r")
	escapedValue = strings.ReplaceAll(escapedValue, "\"", "\\\"")
	setScript := fmt.Sprintf(`window["%s"] = "%s";`,
		strings.TrimLeft(name, "$"), escapedValue,
		)

	var result string
	err := chromedp.Run(browser.chromeContext,
		chromedp.Evaluate(setScript, &result),
	)

	assertErr("SetVariable", err)
}
