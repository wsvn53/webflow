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
