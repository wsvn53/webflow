package main

import (
	"fmt"
	"strings"
)

type FlowImplPrint FlowImplBase

func (impl *FlowImplPrint) Do(args...interface{}) error {
	browser := args[0].(*Browser)

	out := impl.command.Fields[0].ToString()
	if _, ok := browser.variableMaps[out]; ok && strings.HasPrefix(out, "$") {
		fmt.Println(browser.variableMaps[out])
		return nil
	}
	fmt.Println(out)

	return nil
}

//go:generate make IMPL_TYPE=FlowImplPrint gen-impl


func (impl *FlowImplPrint) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplPrint) Type() FlowImplType {
	return impl.implType
}
