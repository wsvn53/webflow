package main

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplSetCookie FlowImplBase

func (impl *FlowImplSetCookie) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplSetCookie) Usage() string {
	return impl.Name() + " <name> <value> <domain> 	Set cookie key/value to browser context with <domain>"
}

func (impl *FlowImplSetCookie) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}

	browser := args[0].(*Browser)
	cookieName := impl.command.GetFieldString(0)
	cookieValue := impl.command.GetFieldString(1)
	cookieDomain := impl.command.GetFieldString(2)

	err := chromedp.Run(browser.chromeContext, chromedp.ActionFunc(func(ctx context.Context) error {
		cookieParam := network.SetCookie(cookieName, cookieValue)
		if cookieDomain != "" {
			cookieParam = cookieParam.WithDomain(cookieDomain)
		}
		_, err := cookieParam.Do(ctx)
		return err
	}))

	return err
}

//go:generate make IMPL_TYPE=FlowImplSetCookie gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplSetCookie{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplSetCookie) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplSetCookie) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplSetCookie) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplSetCookie) Clone() IFlowImpl {
	c := &FlowImplSetCookie{}
	_ = copier.Copy(c, impl)
	return c
}
