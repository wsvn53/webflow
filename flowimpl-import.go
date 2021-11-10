package main

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type FlowImplImport FlowImplBase

func (impl *FlowImplImport) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplImport) Usage() string {
	return impl.Name() + " <script|javascript_file> 	Import <script|javascript_file> to evaluate on new document created"
}

func (impl *FlowImplImport) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	script := impl.command.GetFieldString(0)

	// check script is a file and exists
	if _, err := os.Stat(script); os.IsNotExist(err) == false {
		// read file content to script
		bytes, err := ioutil.ReadFile(script)
		assertErr("Import:", err)
		script = string(bytes)
    }

	err := chromedp.Run(browser.chromeContext, chromedp.ActionFunc(func(ctx context.Context) error {
		_, err := page.AddScriptToEvaluateOnNewDocument(script).Do(ctx)
		if err != nil {
			return err
		}
		return nil
	}))
	return err
}

//go:generate make IMPL_TYPE=FlowImplImport gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplImport{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplImport) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplImport) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplImport) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplImport) Clone() IFlowImpl {
	c := &FlowImplImport{}
	_ = copier.Copy(c, impl)
	return c
}
