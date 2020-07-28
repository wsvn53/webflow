package main

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplVar FlowImplBase

func (impl *FlowImplVar) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplVar) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	variableName := impl.command.Fields[0].ToString()
	variableValue := impl.command.Fields[1].ToString()

	if strings.HasPrefix(variableName, "$") == false {
		return errors.New(fmt.Sprintf("%s is not a valid variable.", variableName))
	}

	browser.SetVariable(variableName, variableValue)
	return nil
}

//go:generate make IMPL_TYPE=FlowImplVar gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplVar{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplVar) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplVar) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplVar) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplVar) Clone() IFlowImpl {
	c := &FlowImplVar{}
	_ = copier.Copy(c, impl)
	return c
}
