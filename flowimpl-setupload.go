package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplSetUpload FlowImplBase

func (impl *FlowImplSetUpload) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplSetUpload) Usage() string {
	return impl.Name() + " <selector> <file1> <file2> ... 	Set files of element <selector> for upload"
}

func (impl *FlowImplSetUpload) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	textSelector := impl.command.GetFieldString(0)
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

//go:generate make IMPL_TYPE=FlowImplSetUpload gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplSetUpload{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplSetUpload) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplSetUpload) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplSetUpload) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplSetUpload) Clone() IFlowImpl {
	c := &FlowImplSetUpload{}
	_ = copier.Copy(c, impl)
	return c
}
