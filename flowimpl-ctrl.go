package main

import (
	"context"
	"strconv"
	"time"
)

/* Impl for all DOM operation commands */

// impl [timeout] command
type FlowImplTimeout FlowImplBase

func (impl *FlowImplTimeout) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	timeout, _ := strconv.Atoi(impl.Command().Fields[0].ToString())
	browser.chromeContext, browser.chromeCancel = context.WithTimeout(
		browser.chromeContext, time.Second * time.Duration(timeout) )
	return nil
}

func (impl *FlowImplTimeout) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplTimeout) Type() FlowImplType {
	return impl.implType
}
