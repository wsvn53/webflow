package main

// impl [eval] command
type FlowImplEval FlowImplBase
func (impl *FlowImplEval) Do() error {
	return nil
}

func (impl *FlowImplEval) Command() *FlowCommand {
	return impl.command
}
