package main

import (
	"context"
	"strconv"
	"time"
)

type FlowImplTimeout FlowImplBase

func (impl *FlowImplTimeout) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	timeout, _ := strconv.Atoi(impl.Command().Fields[0].ToString())
	browser.chromeContext, browser.chromeCancel = context.WithTimeout(
		browser.chromeContext, time.Millisecond * time.Duration(timeout))
	return nil
}

//go:generate make IMPL_TYPE=FlowImplTimeout gen-impl


func (impl *FlowImplTimeout) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplTimeout) Type() FlowImplType {
	return impl.implType
}
