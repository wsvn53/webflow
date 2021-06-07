package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mkideal/cli"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

const VERSION = "v0.4"

type FlowOpts struct {
	Flowfile		string	`cli:"f,file" usage:"Specify Flowfile path."`
	FlowContent		string	`cli:"c,flow" usage:"Using raw flow content string."`
	InsertContent	string	`cli:"i,insert" usage:"Insert new flow before the flow content."`
	AppendContent	string	`cli:"a,append" usage:"Append new flow to the end of flow content."`
	PrintVersion 	bool	`cli:"v,version" usage:"Print webflow version."`
	VerboseMode		bool 	`cli:"d,verbose" usage:"Verbose detail mode."`
	Help			bool	`cli:"h,help" usage:"Show help."`
}

func (flowOpts *FlowOpts) AutoHelp() bool {
	if flowOpts.Help == false {
		return false
	}

	_, _ = color.New(color.Bold).Println("Usage:")

	var flowUsages []string
	for _, flow := range registeredFlows {
		if m, ok := flow.(IFlowUsage); ok {
			flowUsages = append(flowUsages, "  " + m.Usage())
		}
	}

	sort.Strings(flowUsages)

	flowUsages = append(flowUsages, "")
	flowUsages = append(flowUsages, "  [..]	Parameter is optional;")
	flowUsages = append(flowUsages, "  <..>	Parameter is required;")
	flowUsages = append(flowUsages, "  ...	Follow with one or more parameters;")
	flowUsages = append(flowUsages, "\n")
	fmt.Printf("\n%s", strings.Join(flowUsages, "\n"))

	return true
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
		browser.setLogEnable(flowOpts.VerboseMode)
		_ = browser.Run()

		return nil
	}))
}