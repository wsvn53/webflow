package main

import (
	"context"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type FlowImplTimeout FlowImplBase

func (impl *FlowImplTimeout) Type() FlowImplType {
	return FlowImplTypeCtrl
}

func (impl *FlowImplTimeout) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	timeout, _ := strconv.Atoi(impl.Command().Fields[0].ToString())
	browser.chromeContext, browser.chromeCancel = context.WithTimeout(
		browser.chromeContext, time.Millisecond * time.Duration(timeout))
	return nil
}

//go:generate make IMPL_TYPE=FlowImplTimeout gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplTimeout{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplTimeout) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplTimeout) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplTimeout) Command() *FlowCommand {
	return impl.command
}
