package main

import (
	"github.com/chromedp/chromedp"
)

type FlowImplHeadless FlowImplBase

func (impl *FlowImplHeadless) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}

	opt := args[0].(*chromedp.ExecAllocatorOption)
	if len(impl.command.Fields) == 0 {
		*opt = chromedp.Flag("headless", true)
	}

	switch impl.command.Fields[0].ToString() {
	case "true":
		fallthrough
	case "YES":
		fallthrough
	case "1":
		*opt = chromedp.Flag("headless", true)
	case "false":
		fallthrough
	case "NO":
		fallthrough
	case "0":
		*opt = chromedp.Flag("headless", false)
	}

	return nil
}

//go:generate make IMPL_TYPE=FlowImplHeadless gen-impl


func (impl *FlowImplHeadless) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplHeadless) Type() FlowImplType {
	return impl.implType
}
