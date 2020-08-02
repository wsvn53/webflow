package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplFiles FlowImplBase

func (impl FlowImplFiles) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplFiles) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	textSelector := impl.command.Fields[0].ToString()
	var files []string
	for _, f := range impl.command.Fields[1:] {
		fvalue := f.ToString()
		if strings.HasPrefix(fvalue, "$") {
			textValues, ok := browser.variableMaps[fvalue]
			if ok {
				fvalue = textValues
			}
		}
		files = append(files, fvalue)
	}

	err := chromedp.Run(browser.chromeContext,
		chromedp.SetUploadFiles(textSelector, files, chromedp.ByQuery),
	)
	return err
}

//go:generate make IMPL_TYPE=FlowImplFiles gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplFiles{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplFiles) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplFiles) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplFiles) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplFiles) Clone() IFlowImpl {
	c := &FlowImplFiles{}
	_ = copier.Copy(c, impl)
	return c
}
