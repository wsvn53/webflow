package main

import (
	"github.com/cosiner/flag"
	"io/ioutil"
	"os"
)

type Cli struct {
	Flowfile	string	`names:"-f, --file" usage:"Specify Flowfile."`
}

func (t *Cli) Metadata() map[string]flag.Flag {
	return map[string]flag.Flag{
		"": {
			Usage:   "webflow with Flowfile",
			Version: "v0.0.1",
			Desc:    "Webflow can perform a series of web operations defined by Flowfile.",
		},
		"--file": {
			Desc: "Specify Flowfile.",
		},
	}
}

func main() {
	var cli Cli
	err := flag.NewFlagSet(flag.Flag{}).ParseStruct(&cli, os.Args...)
	assertErr("FlagSet", err)

	flowBytes, err := ioutil.ReadFile(cli.Flowfile)
	assertErr("Flowfile", err)
	flowContents := string(flowBytes)

	flow := FlowFromString(flowContents)
	browser := NewBrowser(flow)
	_ = browser.Run()
}