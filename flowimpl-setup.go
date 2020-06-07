package main

import (
	"errors"
	"github.com/chromedp/chromedp"
	"strconv"
	"strings"
)

/* Impl for all chromedp environment Setup operation commands */

// impl [useragent] command
type FlowImplUserAgent FlowImplBase
func (impl *FlowImplUserAgent) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	opt := args[0].(*chromedp.ExecAllocatorOption)
	*opt = chromedp.UserAgent(impl.command.Fields[0].ToString())
	return nil
}

func (impl *FlowImplUserAgent) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplUserAgent) Type() FlowImplType {
	return impl.implType
}

// impl [screen] command
type FlowImplScreen FlowImplBase

func (impl *FlowImplScreen) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}

	opt := args[0].(*chromedp.ExecAllocatorOption)

	if len(impl.command.Fields) == 0 {
		return errors.New("command 'screen' has no arguments")
	}
	sizes := strings.Split(impl.command.Fields[0].ToString(), "x")
	if len(sizes) != 2 {
		return errors.New("illegal screen size")
	}
	w, _ := strconv.Atoi(sizes[0])
	h, _ := strconv.Atoi(sizes[1])
	*opt = chromedp.WindowSize(w, h)

	return nil
}

func (impl *FlowImplScreen) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplScreen) Type() FlowImplType {
	return impl.implType
}


// impl [head] command
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

func (impl *FlowImplHeadless) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplHeadless) Type() FlowImplType {
	return impl.implType
}
