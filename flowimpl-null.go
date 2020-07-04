package main

type FlowImplNull FlowImplBase

func (impl *FlowImplNull) Do(args...interface{}) error {
	// null task do nothing
	return nil
}

//go:generate make IMPL_TYPE=FlowImplNull gen-impl


func (impl *FlowImplNull) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplNull) Type() FlowImplType {
	return impl.implType
}
