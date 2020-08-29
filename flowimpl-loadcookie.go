package main

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
)

type FlowImplLoadCookie FlowImplBase

func (impl *FlowImplLoadCookie) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplLoadCookie) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}

	browser := args[0].(*Browser)
	cookieFile := impl.command.GetFieldString(0)
	cookiePairs, err := ioutil.ReadFile(cookieFile)
	if err != nil{
		return err
	}
	cookieDomain := impl.command.GetFieldString(1)

	cookieRegexp := regexp.MustCompile("([^= ]+)=([^; ]*)[;]?")
	cookieMatches := cookieRegexp.FindAllStringSubmatch(string(cookiePairs), -1)
	err = chromedp.Run(browser.chromeContext, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		for _, c := range cookieMatches {
			param := network.SetCookie(c[1], c[2])
			if browser.logFunc != nil {
				_, _ = (*browser.logFunc)("-", param.Name, "=", param.Value)
			}
			if cookieDomain != "" {
				param = param.WithDomain(cookieDomain)
			}
			_, err = param.Do(ctx)
			assertErr("LoadCookie", err)
		}
		return err
	}))

	return err
}

//go:generate make IMPL_TYPE=FlowImplLoadCookie gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplLoadCookie{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplLoadCookie) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplLoadCookie) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplLoadCookie) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplLoadCookie) Clone() IFlowImpl {
	c := &FlowImplLoadCookie{}
	_ = copier.Copy(c, impl)
	return c
}
