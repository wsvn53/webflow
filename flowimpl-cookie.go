package main

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplCookie FlowImplBase

func (impl *FlowImplCookie) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplCookie) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}

	browser := args[0].(*Browser)
	cookieName := impl.command.GetFieldString(0)
	cookieValue := impl.command.GetFieldString(1)
	cookieDomain := impl.command.GetFieldString(2)

	err := chromedp.Run(browser.chromeContext, chromedp.ActionFunc(func(ctx context.Context) error {
		_, err := network.SetCookie(cookieName, cookieValue).
			WithDomain(cookieDomain).Do(ctx)
		return err
	}))

	return err
}

//go:generate make IMPL_TYPE=FlowImplCookie gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplCookie{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplCookie) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplCookie) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplCookie) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplCookie) Clone() IFlowImpl {
	c := &FlowImplCookie{}
	_ = copier.Copy(c, impl)
	return c
}
