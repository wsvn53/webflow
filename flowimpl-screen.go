package main

import (
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"reflect"
	"strconv"
	"strings"
)

type FlowImplScreen FlowImplBase

func (impl *FlowImplScreen) Type() FlowImplType {
	return FlowImplTypeFlag
}

func (impl *FlowImplScreen) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	opt := args[0].(*chromedp.ExecAllocatorOption)
	sizes := strings.Split(impl.command.GetFieldString(0), "x")
	if len(sizes) != 2 {
		return nil
	}
	w, _ := strconv.Atoi(sizes[0])
	h, _ := strconv.Atoi(sizes[1])
	*opt = chromedp.WindowSize(w, h)
	return nil
}

//go:generate make IMPL_TYPE=FlowImplScreen gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplScreen{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplScreen) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplScreen) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplScreen) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplScreen) Clone() IFlowImpl {
	c := &FlowImplScreen{}
	_ = copier.Copy(c, impl)
	return c
}
