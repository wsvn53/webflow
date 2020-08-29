package main

import (
	"github.com/jinzhu/copier"
	"os/exec"
	"reflect"
	"strings"
)

type FlowImplShell FlowImplBase

func (impl *FlowImplShell) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplShell) Usage() string {
	return impl.Name() + " <script|$variable> 	Run custom shell scripts, support $variable"
}

func (impl *FlowImplShell) Do(args...interface{}) error {
	browser := args[0].(*Browser)
	shellScript := impl.command.GetFieldString(0)
	var variableName string
	if len(impl.command.Fields) > 1 {
		variableName = impl.command.GetFieldString(0)
		shellScript = impl.command.GetFieldString(1)
	}

	shellCommand := exec.Command("sh", "-c", shellScript)
	shellResult, _ := shellCommand.CombinedOutput()

	if strings.HasPrefix(variableName, "$") {
		browser.SetVariable(variableName, string(shellResult))
	}

	return nil
}

//go:generate make IMPL_TYPE=FlowImplShell gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplShell{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplShell) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplShell) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplShell) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplShell) Clone() IFlowImpl {
	c := &FlowImplShell{}
	_ = copier.Copy(c, impl)
	return c
}
