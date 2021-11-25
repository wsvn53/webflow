package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type FlowImplWait FlowImplBase

func (impl *FlowImplWait) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplWait) Usage() string {
	return impl.Name() + " <selector|milliseconds> 	Wait for target with <selector> OR timeout with <milliseconds>"
}

func (impl *FlowImplWait) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	waitTarget := impl.command.GetFieldString(0)
	timeout, err := strconv.Atoi(waitTarget)
	if err == nil {
		// wait for timeout
		err = chromedp.Run(browser.chromeContext,
			chromedp.Sleep(time.Millisecond * time.Duration(timeout)))
		return err
	}

	err = chromedp.Run(browser.chromeContext,
		chromedp.WaitVisible(waitTarget, chromedp.ByQuery, chromedp.FromNode(browser.switchNode)),
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

func (impl *FlowImplWait) Clone() IFlowImpl {
	c := &FlowImplWait{}
	_ = copier.Copy(c, impl)
	return c
}
