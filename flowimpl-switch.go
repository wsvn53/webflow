package main

import (
	"errors"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplSwitch FlowImplBase

func (impl FlowImplSwitch) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplSwitch) Usage() string {
	return impl.Name() + " <URL> 	Switch context which URL contains <URL>."
}

func (impl *FlowImplSwitch) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	url := impl.command.GetFieldString(0)

	if url == "" {
		browser.switchNode = nil
		return nil
	}

	var iframes []*cdp.Node
	err := chromedp.Run(browser.chromeContext, chromedp.Nodes(`iframe`, &iframes, chromedp.ByQueryAll))
	if err != nil {
		return err
	}

	if len(iframes) == 0 {
		return errors.New("no iframe context found")
	}

	for _, iframe := range iframes {
		if strings.Contains(iframe.AttributeValue("src"), url) {
			browser.switchNode = iframe
			break
		}
	}

	if browser.switchNode == nil {
		return errors.New("cannot find iframe context matched <" + url + ">")
	}

	return nil
}

//go:generate make IMPL_TYPE=FlowImplSwitch gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplSwitch{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplSwitch) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplSwitch) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplSwitch) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplSwitch) Clone() IFlowImpl {
	c := &FlowImplSwitch{}
	_ = copier.Copy(c, impl)
	return c
}
