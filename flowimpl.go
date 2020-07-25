package main

import "strings"

type FlowImplBase struct {
	implType 		FlowImplType
	command 		*FlowCommand
}

type IFlowImpl interface {
	Do(args...interface{})		error
	Name()		string
	Command()	*FlowCommand
	Type()		FlowImplType
	SetCommand(command *FlowCommand)
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

var registeredFlows map[string]IFlowImpl

func registerFlow(flowImpl IFlowImpl) {
	if registeredFlows == nil {
		registeredFlows = map[string]IFlowImpl{}
	}
	registeredFlows[flowImpl.Name()] = flowImpl
}

func NewFlowImpl(command *FlowCommand) IFlowImpl {
	if flowImpl, ok := registeredFlows[strings.ToLower(*command.Name)]; ok {
		flowImpl.SetCommand(command)
		return flowImpl
	}
	return &FlowImplNull{}
}
