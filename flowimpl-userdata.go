package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/go-homedir"
	"reflect"
	"strings"
)

type FlowImplUserData FlowImplBase

func (impl *FlowImplUserData) Type() FlowImplType {
	return FlowImplTypeFlag
}

func (impl *FlowImplUserData) Usage() string {
	return impl.Name() + " <string> 	Setup browser userdata storage path."
}

func (impl *FlowImplUserData) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	opt := args[0].(*chromedp.ExecAllocatorOption)
	userDir := impl.command.GetFieldString(0)
	userDir, _ = homedir.Expand(userDir)
	*opt = chromedp.Flag("user-data-dir", userDir)
	return nil
}

//go:generate make IMPL_TYPE=FlowImplUserData gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplUserData{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplUserData) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplUserData) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplUserData) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplUserData) Clone() IFlowImpl {
	c := &FlowImplUserData{}
	_ = copier.Copy(c, impl)
	return c
}
