package main

import "strings"

type FlowImplBase struct {
	implType 		FlowImplType
	command 		*FlowCommand
}

type IFlowImpl interface {
	Do(args...interface{})		error
	Command()	*FlowCommand
	Type()		FlowImplType
}

// flow impl type
type FlowImplType int
const (
	FlowImplTypeFlag     = 0
	FlowImplTypeCtrl     = iota
	FlowImplTypeNavigate = iota
	FlowImplTypeDom      = iota
	FlowImplTypeJsop     = iota
)

func NewFlowImpl(command *FlowCommand) IFlowImpl {
	switch strings.ToLower(*command.Name) {
	case "open":
		return &FlowImplOpen{ command: command, implType: FlowImplTypeNavigate }
	case "text":
		return &FlowImplText{ command: command, implType: FlowImplTypeDom }
	case "eval":
		return &FlowImplEval{ command: command, implType: FlowImplTypeJsop }
	case "timeout":
		return &FlowImplTimeout{ command: command, implType: FlowImplTypeCtrl }
	case "useragent":
		return &FlowImplUserAgent{ command: command, implType: FlowImplTypeFlag}
	case "screen":
		return &FlowImplScreen{ command: command, implType: FlowImplTypeFlag}
	case "headless":
		return &FlowImplHeadless{ command: command, implType: FlowImplTypeFlag}
	}
	return nil
}
