package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"io/ioutil"
	"os"
)

const VERSION = "v0.0.1"

type FlowOpts struct {
	Flowfile	string	`cli:"f,file" usage:"Specify Flowfile path."`
	FlowContent	string	`cli:"c,flow" usage:"Using raw flow content string."`
	InsertContent	string	`cli:"i,insert" usage:"Insert new flow before the flow content."`
	AppendContent	string	`cli:"a,append" usage:"Append new flow to the end of flow content."`
	PrintVersion 	bool	`cli:"version" usage:"Print webflow version."`
	VerboseMode		bool 	`cli:"v,verbose" usage:"Verbose mode."`
	Help		bool	`cli:"h,help" usage:"Show help."`
}

func (flowOpts *FlowOpts) AutoHelp() bool {
	return flowOpts.Help
}

func main() {
	os.Exit(cli.Run(new(FlowOpts), func(context *cli.Context) error {
		flowOpts := context.Argv().(*FlowOpts)

		if flowOpts.PrintVersion {
			fmt.Println(os.Args[0], VERSION)
			os.Exit(0)
		}

		flowContents := flowOpts.FlowContent

		if len(flowOpts.Flowfile) > 0 {
			flowBytes, err := ioutil.ReadFile(flowOpts.Flowfile)
			assertErr("Flowfile", err)
			flowContents = string(flowBytes)
		}

		if len(flowContents) == 0 {
			flowBytes, err := ioutil.ReadAll(os.Stdin)
			assertErr("Stdin", err)
			flowContents = string(flowBytes)
			fmt.Println(flowContents)
		}

		// --insert new content
		if len(flowOpts.InsertContent) > 0 {
			flowContents = flowOpts.InsertContent + "\n" + flowContents
		}

		// --append new content
		if len(flowOpts.AppendContent) > 0 {
			flowContents = flowContents + "\n" + flowOpts.AppendContent
		}

		flow := FlowFromString(flowContents)
		browser := NewBrowser(flow)
		if flowOpts.VerboseMode {
			logFunc := func(a ...interface{}) (n int, err error) {
				return fmt.Fprintln(os.Stderr, a...)
			}
			browser.logFunc = &logFunc
		}
		_ = browser.Run()

		return nil
	}))
}