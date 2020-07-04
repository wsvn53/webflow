package main

import (
	"github.com/chromedp/chromedp"
)

type FlowImplUserAgent FlowImplBase
func (impl *FlowImplUserAgent) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	opt := args[0].(*chromedp.ExecAllocatorOption)
	*opt = chromedp.UserAgent(impl.command.Fields[0].ToString())
	return nil
}

//go:generate make IMPL_TYPE=FlowImplUserAgent gen-impl


func (impl *FlowImplUserAgent) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplUserAgent) Type() FlowImplType {
	return impl.implType
}
