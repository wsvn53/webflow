package main

/* Impl for all DOM operation commands */

// impl [timeout] command
type FlowImplTimeout FlowImplBase

func (impl *FlowImplTimeout) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	return nil
}

func (impl *FlowImplTimeout) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplTimeout) Type() FlowImplType {
	return impl.implType
}
