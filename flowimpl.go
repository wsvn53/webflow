package main

import "strings"

type FlowImplBase struct {
	command 	*FlowCommand
}

type IFlowImpl interface {
	Do()		error
	Command()	*FlowCommand
}

// impl [open] command
type FlowImplOpen FlowImplBase

func (impl *FlowImplOpen) Do() error {
	return nil
}

func (impl *FlowImplOpen) Command() *FlowCommand {
	return impl.command
}

// impl [text] command
type FlowImplText FlowImplBase
func (impl *FlowImplText) Do() error {
	return nil
}

func (impl *FlowImplText) Command() *FlowCommand {
	return impl.command
}

// impl [eval] command
type FlowImplEval FlowImplBase
func (impl *FlowImplEval) Do() error {
	return nil
}

func (impl *FlowImplEval) Command() *FlowCommand {
	return impl.command
}

// impl [timeout] command
type FlowImplTimeout FlowImplBase
func (impl *FlowImplTimeout) Do() error {
	return nil
}

func (impl *FlowImplTimeout) Command() *FlowCommand {
	return impl.command
}

// impl [useragent] command
type FlowImplUserAgent FlowImplBase
func (impl *FlowImplUserAgent) Do() error {
	return nil
}

func (impl *FlowImplUserAgent) Command() *FlowCommand {
	return impl.command
}

func NewFlowImpl(command *FlowCommand) IFlowImpl {
	switch strings.ToLower(*command.Name) {
	case "open":
		return &FlowImplOpen{ command: command }
	case "text":
		return &FlowImplText{ command: command }
	case "eval":
		return &FlowImplEval{ command: command }
	case "timeout":
		return &FlowImplTimeout{ command: command }
	case "useragent":
		fallthrough
	case "user-agent":
		return &FlowImplUserAgent{ command: command }
	}
	return nil
}
