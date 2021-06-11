package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Browser struct {
	webFlow 		*Flow
	switchNode 		*cdp.Node
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
	defer browser.Cancel()

	// enable keyboard interactive
	go func() { _ = browser.keyboardLoop() } ()

	browser.webFlow.Walk(func(i int, impl IFlowImpl, stop *bool) {
		if impl == nil || impl.Type() == FlowImplTypeFlag || impl.Type() == FlowImplTypeLog {
			return
		}
		if browser.logFunc != nil {
			if reflect.TypeOf(impl) == reflect.TypeOf(&FlowImplNull{}) {
				browser.Println("> [!] Invalid:", *impl.Command().Name, impl.Command().FieldsString())
				return
			}
			browser.Println("> Task:", *impl.Command().Name, impl.Command().FieldsString())
		}
		err := impl.Do(browser)
		if err != nil {
			browser.Cancel()
		}
		assertErr("Run", err)
	})

	return nil
}

func (browser *Browser) Cancel() {
	_ = keyboard.Close()
	if browser.chromeContext != nil {
		chromedp.Cancel(browser.chromeContext)
	}
}

func (browser *Browser) SetVariable(name string, value interface{}) {
	valueText := fmt.Sprintf("%v", value)
	browser.variableMaps[name] = valueText

	// store variable to web context
	escapedValue := strconv.Quote(valueText)
	setScript := fmt.Sprintf(`window["%s"] = %s;`, name, escapedValue)

	var result interface{}
	err := chromedp.Run(browser.chromeContext,
		chromedp.Evaluate(setScript, &result),
	)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Script:", setScript)
	}
	assertErr("SetVariable", err)
}

func (browser *Browser) setLogEnable(enable bool) {
	if enable == false {
		browser.logFunc = nil
	}

	logFunc := func(a ...interface{}) (n int, err error) {
		if strings.HasPrefix(a[0].(string), "> ") {
			return color.New(color.FgBlue).Fprintln(os.Stderr, a...)
		}
		return fmt.Fprintln(os.Stderr, a...)
	}
	browser.logFunc = &logFunc
}

func (browser *Browser) Println(a ...interface{}) {
	_, _ = (*browser.logFunc)(a...)
}

func (browser *Browser) keyboardLoop() error {
	defer browser.Cancel()

	keyEvents, err := keyboard.GetKeys(2)
	if err != nil { return err }

	browser.Println("> Using CTRL+S to Take Screenshot..")
	for {
		event := <- keyEvents
		assertErr("Key", event.Err)
		switch event.Key {
		case keyboard.KeyCtrlC:
			browser.Cancel()
			os.Exit(0)
		case keyboard.KeyCtrlS:
			screenshotImpl := new(FlowImplScreenshot)
			screenPath := filepath.Join(".", fmt.Sprintf("Screen-%d.png", time.Now().UnixNano()))
			browser.Println("- Screenshot:", screenPath)
			err := screenshotImpl.takeScreenshot(browser, "", screenPath)
			if err != nil {
				browser.Println("[!] Screenshot:", err)
			}
		}
	}
}
