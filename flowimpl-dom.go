package main

// impl [text] command
type FlowImplText FlowImplBase
func (impl *FlowImplText) Do() error {
	return nil
}

func (impl *FlowImplText) Command() *FlowCommand {
	return impl.command
}
