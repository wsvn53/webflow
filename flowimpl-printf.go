package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type FlowImplPrintf FlowImplBase

func (impl *FlowImplPrintf) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplPrintf) Usage() string {
	return impl.Name() + " <format> <field1> <$variable> ... 	Print format string, bridge to fmt.Printf function"
}

func (impl *FlowImplPrintf) Do(args...interface{}) error {
	browser := args[0].(*Browser)

	format := impl.command.GetFieldString(0)
	var fields []interface{}
	for _, v := range impl.command.Fields[1:] {
		vs := v.ToString()
		if value, ok := browser.variableMaps[vs]; ok && strings.HasPrefix(vs, "$") {
			fields = append(fields, value)
			continue
		}
		fields = append(fields, vs)
	}

	if len(fields) == 0 {
		return nil
	}

	fmt.Printf(format, fields...)
	return nil
}

//go:generate make IMPL_TYPE=FlowImplPrintf gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplPrintf{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplPrintf) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplPrintf) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplPrintf) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplPrintf) Clone() IFlowImpl {
	c := &FlowImplPrintf{}
	_ = copier.Copy(c, impl)
	return c
}
