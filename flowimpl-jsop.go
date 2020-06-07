package main

/* Impl for all Javascript operation commands */

// impl [eval] command
type FlowImplEval FlowImplBase

func (impl *FlowImplEval) Do(args...interface{}) error {
	return nil
}

func (impl *FlowImplEval) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplEval) Type() FlowImplType {
	return impl.implType
}
