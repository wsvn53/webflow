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
	FlowImplTypeOP       = iota		// common operation
	FlowImplTypeLog      = iota
	FlowImplTypeNull     = iota		// unknown task
)

func NewFlowImpl(command *FlowCommand) IFlowImpl {
	switch strings.ToLower(*command.Name) {
	case "timeout":
		return &FlowImplTimeout{ command: command, implType: FlowImplTypeCtrl }
	case "open":
		return &FlowImplOpen{ command: command, implType: FlowImplTypeOP }
	case "value":
		return &FlowImplValue{ command: command, implType: FlowImplTypeOP }
	case "keys":
		return &FlowImplKeys{ command: command, implType: FlowImplTypeOP }
	case "click":
		return &FlowImplClick{ command: command, implType: FlowImplTypeOP }
	case "eval":
		return &FlowImplEval{ command: command, implType: FlowImplTypeOP}
	case "print":
		return &FlowImplPrint{ command: command, implType: FlowImplTypeOP }
	case "log":
		return &FlowImplLog{ command: command, implType: FlowImplTypeOP }
	case "wait":
		return &FlowImplWait{ command: command, implType: FlowImplTypeOP }
	case "useragent":
		return &FlowImplUserAgent{ command: command, implType: FlowImplTypeFlag }
	case "screen":
		return &FlowImplScreen{ command: command, implType: FlowImplTypeFlag }
	case "headless":
		return &FlowImplHeadless{ command: command, implType: FlowImplTypeFlag }
	case "debug":
		return &FlowImplDebug{ command: command, implType: FlowImplTypeLog }
	default:
		return &FlowImplNull{ command: command, implType: FlowImplTypeNull }
	}
}
