package main

import (
	"github.com/chromedp/chromedp"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type FlowImplWait FlowImplBase

func (impl *FlowImplWait) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplWait) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	waitTarget := impl.command.Fields[0].ToString()
	timeout, err := strconv.Atoi(waitTarget)
	if err == nil {
		// wait for timeout
		err = chromedp.Run(browser.chromeContext,
			chromedp.Sleep(time.Millisecond * time.Duration(timeout)))
		return err
	}

	err = chromedp.Run(browser.chromeContext,
		chromedp.WaitReady(waitTarget, chromedp.ByQuery),
	)
	return err
}

//go:generate make IMPL_TYPE=FlowImplWait gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplWait{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplWait) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplWait) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplWait) Command() *FlowCommand {
	return impl.command
}
