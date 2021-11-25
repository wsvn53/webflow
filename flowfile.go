package main

import (
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"strconv"
	"strings"
)

type Flow struct {
	contents 	string
	FlowCommands	[]*FlowCommand		`{ @@ }`

	flowImpls 		[]IFlowImpl
}

type FlowCommand struct {
	Pos		lexer.Position

	Name 	*string			`@Ident`
	Fields	[]FlowField		`{ @@ } [ ";" ]`
}

type FlowField struct {
	Variable	*string		`"$" @(Ident)`
	String 		*string 	`| @(Ident | String | RawString)`
	Int			*int		`| @Int`
}

func (command *FlowCommand) GetFieldString(i int) string {
	if i < len(command.Fields) {
		return command.Fields[i].ToString()
	}
	return ""
}

func (command *FlowCommand) GetFieldInt(i int) int {
	var ret int
	if i < len(command.Fields) {
		ret, _ = strconv.Atoi(command.Fields[i].ToString())
	}
	return ret
}

func (command *FlowCommand) FieldsString() string {
	var fs []string
	for _, field := range command.Fields {
		fs = append(fs, field.ToString())
	}
	return fmt.Sprintf("[%s]", strings.Join(fs, "], ["))
}

func (field *FlowField) ToString() string {
	switch {
	case field.isVariable():
		return "$" + *field.Variable
	case field.String != nil:
		return *field.String
	case field.Int != nil:
		return fmt.Sprintf("%d", *field.Int)
	}
	return ""
}

func (field *FlowField) isVariable() bool {
	return field.Variable != nil
}

func FlowFromString(flowContents string) *Flow {
	parser, err := participle.Build(&Flow{},
		participle.Elide("Comment"))
	assertErr("FLOWFILE", err)

	flow := &Flow{}
	err = parser.ParseString(flowContents, flow)
	assertErr("FLOWFILE", err)
	flow.contents = flowContents

	return flow
}

func (flow *Flow) CommandByName(name string) *FlowCommand {
	if flow.FlowCommands == nil {
		return nil
	}

	for _, cmd := range flow.FlowCommands {
		if *cmd.Name == name {
			return cmd
		}
	}

	return nil
}

func (flow *Flow) Impls() []IFlowImpl {
	if len(flow.flowImpls) > 0 {
		return flow.flowImpls
	}

	for _, cmd := range flow.FlowCommands {
		flow.flowImpls = append(flow.flowImpls, NewFlowImpl(cmd))
	}

	return flow.flowImpls
}

func (flow *Flow) Walk(walkFunc func(i int, impl IFlowImpl, stop *bool)) {
	if flow.Impls() == nil {
		return
	}

	var stop bool
	for i, impl := range flow.Impls() {
		walkFunc(i, impl, &stop)
		if stop == true {
			return
		}
	}
}

func (flow *Flow) WalkByType(implType FlowImplType, walkFunc func(i int, impl IFlowImpl, stop *bool)) {
	flow.Walk(func(i int, impl IFlowImpl, stop *bool) {
		if impl == nil || impl.Type() != implType {
			return
		}
		walkFunc(i, impl, stop)
	})
}
