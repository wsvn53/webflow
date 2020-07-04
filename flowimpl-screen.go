package main

type FlowImplScreen FlowImplBase
func (impl *FlowImplScreen) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	return nil
}

//go:generate make IMPL_TYPE=FlowImplScreen gen-impl


func (impl *FlowImplScreen) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplScreen) Type() FlowImplType {
	return impl.implType
}
