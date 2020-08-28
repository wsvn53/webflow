package main

import (
	"github.com/jinzhu/copier"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type FlowImplSave FlowImplBase

func (impl *FlowImplSave) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplSave) Do(args...interface{}) error {
	browser := args[0].(*Browser)

	path := impl.command.GetFieldString(0)
	content := impl.command.GetFieldString(1)
	if _, ok := browser.variableMaps[content]; ok && strings.HasPrefix(content, "$") {
		content = browser.variableMaps[content]
	}
	err := ioutil.WriteFile(path, []byte(content), os.ModePerm)

	return err
}

//go:generate make IMPL_TYPE=FlowImplSave gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplSave{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplSave) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplSave) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplSave) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplSave) Clone() IFlowImpl {
	c := &FlowImplSave{}
	_ = copier.Copy(c, impl)
	return c
}
