package main

import (
	"github.com/chromedp/chromedp"
	"log"
)

type FlowImplDebug FlowImplBase

func (impl *FlowImplDebug) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	opt := args[0].(*chromedp.ContextOption)

	if impl.command.Fields[0].ToString() == "true" {
		*opt = chromedp.WithDebugf(log.Printf)
	}

	return nil
}

//go:generate make IMPL_TYPE=FlowImplDebug gen-impl


func (impl *FlowImplDebug) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplDebug) Type() FlowImplType {
	return impl.implType
}
