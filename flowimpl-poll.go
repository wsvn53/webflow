package main

import (
	"errors"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
	"time"
)

type FlowImplPoll FlowImplBase

func (impl *FlowImplPoll) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplPoll) Usage() string {
	return impl.Name() + "<expressions>	[milliseconds] 	Poll result of <expressions>, timeout with <milliseconds> is optional"
}

func (impl *FlowImplPoll) Do(args...interface{}) error {
	browser := args[0].(*Browser)

	var expr string
	var pollOpts []chromedp.PollOption
	switch {
	case len(impl.command.Fields) == 0:
		return errors.New("no expressions to poll")
	case len(impl.command.Fields) == 2:
		timeout := impl.command.GetFieldInt(1)
		if timeout >= 0 {
			pollOpts = append(pollOpts, chromedp.WithPollingTimeout(time.Duration(timeout) * time.Millisecond))
		}
	}

	expr = impl.command.GetFieldString(0)
	if browser.switchNode != nil {
		pollOpts = append(pollOpts, chromedp.WithPollingInFrame(browser.switchNode))
	}
	err := chromedp.Run(browser.chromeContext,
		chromedp.Poll(expr, nil, pollOpts...),
	)

	// timeout error is not an error
	if err != nil && strings.Contains(err.Error(), "timeout") {
		fmt.Println("timeout")
		return nil
	}

	return err
}

//go:generate make IMPL_TYPE=FlowImplPoll gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplPoll{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplPoll) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplPoll) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplPoll) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplPoll) Clone() IFlowImpl {
	c := &FlowImplPoll{}
	_ = copier.Copy(c, impl)
	return c
}
