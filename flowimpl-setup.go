package main

/* Impl for all chromedp environment Setup operation commands */

// impl [timeout] command
type FlowImplTimeout FlowImplBase
func (impl *FlowImplTimeout) Do() error {
	return nil
}

func (impl *FlowImplTimeout) Command() *FlowCommand {
	return impl.command
}

// impl [useragent] command
type FlowImplUserAgent FlowImplBase
func (impl *FlowImplUserAgent) Do() error {
	return nil
}

func (impl *FlowImplUserAgent) Command() *FlowCommand {
	return impl.command
}
