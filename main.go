package main

import (
	"fmt"
	"github.com/cosiner/flag"
	"io/ioutil"
	"os"
)

const VERSION = "v0.0.1"

type Cli struct {
	Flowfile	string	`names:"-f, --file" usage:"Specify Flowfile path."`
	FlowContent	string	`names:"-c, --flow" usage:"Using raw flow content string."`
	PrintVersion bool	`names:"-v, --version" usage:"Print webflow version."`
}

func (t *Cli) Metadata() map[string]flag.Flag {
	return map[string]flag.Flag{
		"": {
			Usage:   "webflow",
			Version: VERSION,
			Desc:    "Webflow can perform a series of web operations defined by Flowfile.",
		},
		"--file": {
			Desc: "path of flowfile to run.",
		},
		"--flow": {
			Desc: "using flow raw content string.",
		},
	}
}

func main() {
	var cli Cli
	err := flag.NewFlagSet(flag.Flag{}).ParseStruct(&cli, os.Args...)
	assertErr("FlagSet", err)

	if cli.PrintVersion {
		fmt.Println(os.Args[0], VERSION)
		os.Exit(0)
	}

	flowContents := cli.FlowContent

	if len(cli.Flowfile) > 0 {
		flowBytes, err := ioutil.ReadFile(cli.Flowfile)
		assertErr("Flowfile", err)
		flowContents = string(flowBytes)
	}

	if len(flowContents) == 0 {
		flowBytes, err := ioutil.ReadAll(os.Stdin)
		assertErr("Stdin", err)
		flowContents = string(flowBytes)
		fmt.Println(flowContents)
	}

	flow := FlowFromString(flowContents)
	browser := NewBrowser(flow)
	_ = browser.Run()
}