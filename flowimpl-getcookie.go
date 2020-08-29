package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplGetCookie FlowImplBase

func (impl *FlowImplGetCookie) Type() FlowImplType {
	return FlowImplTypeOP
}

func matchCookie(target string, current string) bool {
	if target == "" {
		return true
	}
	return target == current
}

func (impl *FlowImplGetCookie) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}

	browser := args[0].(*Browser)
	varName := impl.command.GetFieldString(0)
	cookieDomain := impl.command.GetFieldString(1)
	cookieName := impl.command.GetFieldString(2)

	cookieOutput := ""
	err := chromedp.Run(browser.chromeContext, chromedp.ActionFunc(func(ctx context.Context) error {
		cookies, err := network.GetCookies().Do(ctx)
		for _, c := range cookies {
			if matchCookie(cookieDomain, c.Domain) && matchCookie(cookieName, c.Name) {
				cookieOutput += fmt.Sprintf("%s=%s; ", c.Name, c.Value)
			}
		}
		return err
	}))
	browser.SetVariable(varName, cookieOutput)
	return err
}

//go:generate make IMPL_TYPE=FlowImplGetCookie gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplGetCookie{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplGetCookie) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplGetCookie) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplGetCookie) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplGetCookie) Clone() IFlowImpl {
	c := &FlowImplGetCookie{}
	_ = copier.Copy(c, impl)
	return c
}
