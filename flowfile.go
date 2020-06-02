package main

import (
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

type Flow struct {
	contents 	string
	FlowCommands	[]*FlowCommand		`{ @@ }`
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