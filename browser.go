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
	logFunc			*func (a ...interface{})(n int, err error)
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
	logOptions := []chromedp.ContextOption{
		chromedp.WithLogf(log.Printf),
		chromedp.WithErrorf(log.Printf),
	}
	flow.WalkByType(FlowImplTypeLog, func(i int, impl IFlowImpl, stop *bool) {
		var opt chromedp.ContextOption
		if _ = impl.Do(&opt); opt != nil {
			logOptions = append(logOptions, opt)
		}
	})
	ctx, cancel = chromedp.NewContext(ctx, logOptions...)

	return &Browser{
		webFlow: flow,
		chromeContext: ctx,
		chromeCancel: cancel,
		variableMaps: map[string]string{},
	}
}

func (browser *Browser) Run() error {
	browser.webFlow.Walk(func(i int, impl IFlowImpl, stop *bool) {
		if impl == nil || impl.Type() == FlowImplTypeFlag || impl.Type() == FlowImplTypeLog {
			return
		}
		if browser.logFunc != nil {
			_, _ = (*browser.logFunc)("> Task:",
				*impl.Command().Name, impl.Command().FieldsString(),
				)
		}
		err := impl.Do(browser)
		assertErr("Run", err)
	})
	return nil
}

func (browser *Browser) SetVariable(name string, value interface{}) {
	valueText := fmt.Sprintf("%v", value)
	browser.variableMaps[name] = valueText

	// store variable to web context
	escapedValue := strings.ReplaceAll(valueText, "\n", "\\n")
	escapedValue = strings.ReplaceAll(escapedValue, "\r", "\\r")
	escapedValue = strings.ReplaceAll(escapedValue, "\"", "\\\"")
	setScript := fmt.Sprintf(`window["%s"] = "%s";`, name, escapedValue)

	var result interface{}
	err := chromedp.Run(browser.chromeContext,
		chromedp.Evaluate(setScript, &result),
	)

	assertErr("SetVariable", err)
}
