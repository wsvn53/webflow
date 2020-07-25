package main

import (
	"fmt"
	"github.com/cosiner/flag"
	"io/ioutil"
	"os"
)

type Cli struct {
	Flowfile	string	`names:"-f, --file" usage:"Specify Flowfile."`
	FlowContent	string	`names:"-c, --flow" usage:"Raw flow contents."`
}

func (t *Cli) Metadata() map[string]flag.Flag {
	return map[string]flag.Flag{
		"": {
			Usage:   "webflow",
			Version: "v0.0.1",
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