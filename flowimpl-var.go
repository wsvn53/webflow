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

func (impl *FlowImplVar) Usage() string {
	return impl.Name() + " $variable <string> 	Define $variable with value <string>"
}

func (impl *FlowImplVar) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	variableName := impl.command.GetFieldString(0)
	variableValue := impl.command.GetFieldString(1)

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
