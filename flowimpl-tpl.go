// +build ignore

package main

func (impl *FlowImplBase) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplBase) Type() FlowImplType {
	return impl.implType
}
