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

type FlowImplGetcookie FlowImplBase

func (impl *FlowImplGetcookie) Type() FlowImplType {
	return FlowImplTypeOP
}

func matchCookie(target string, current string) bool {
	if target == "" {
		return true
	}
	return target == current
}

func (impl *FlowImplGetcookie) Do(args...interface{}) error {
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

//go:generate make IMPL_TYPE=FlowImplGetcookie gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplGetcookie{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplGetcookie) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplGetcookie) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplGetcookie) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplGetcookie) Clone() IFlowImpl {
	c := &FlowImplGetcookie{}
	_ = copier.Copy(c, impl)
	return c
}
