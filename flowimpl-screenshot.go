package main

import (
	"errors"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/copier"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type FlowImplScreenshot FlowImplBase

func (impl *FlowImplScreenshot) Type() FlowImplType {
	return FlowImplTypeOP
}

func (impl *FlowImplScreenshot) Usage() string {
	return impl.Name() + " <selector> <path>  [interval]	Take screenshot of <selector> and save to <path>, [interval] is optional"
}

func (impl *FlowImplScreenshot) Do(args...interface{}) error {
	if len(args) == 0 {
		return nil
	}
	browser := args[0].(*Browser)
	selector := impl.command.GetFieldString(0)
	savepath := impl.command.GetFieldString(1)
	interval := impl.command.GetFieldString(2)

	// satisfy savepath
	fi, err := os.Stat(savepath)
	screename := regexp.MustCompile("[^a-zA-Z0-9_-]").ReplaceAllString(selector, "-")
	if len(screename) == 0 {
		screename = "full"
	}
	switch {
	case fi != nil && fi.IsDir():
		savepath = filepath.Join(savepath, screename + ".png")
	case os.IsNotExist(err) && strings.HasSuffix(savepath, ".png"):
		dirpath := filepath.Dir(savepath)
		_ = os.MkdirAll(dirpath, 0777)
	case os.IsNotExist(err):
		_ = os.MkdirAll(savepath, 0777)
		savepath = filepath.Join(savepath, screename + ".png")
	}

	if interval == "" {
		err = impl.takeScreenshot(browser, selector, savepath)
		return err
	}
	
	ti, err := strconv.Atoi(interval)
	if err != nil {
		return err
	}
	if ti <= 0 {
		return errors.New("[interval] can't be minus number or zero")
	}

	go func() {
		for {
			timer := time.NewTimer(time.Duration(ti) * time.Millisecond)
			<- timer.C
			nameSuffix := fmt.Sprintf("-%d.png", time.Now().UnixNano()/1000)
			sp := regexp.MustCompile(`\.png$`).ReplaceAllString(savepath, nameSuffix)
			err = impl.takeScreenshot(browser, selector, sp)
			_, _ = fmt.Fprintln(os.Stderr, "- Screenshot:", selector, sp)
		}
	}()

	return nil
}

func (impl *FlowImplScreenshot) takeScreenshot(browser *Browser, selector string, savepath string) error {
	var screenBuffer []byte
	var err error
	if selector == "" {
		err = chromedp.Run(browser.chromeContext, chromedp.FullScreenshot(&screenBuffer, 100))
	} else {
		err = chromedp.Run(browser.chromeContext,
			chromedp.Screenshot(selector, &screenBuffer, chromedp.ByQuery, chromedp.NodeVisible, chromedp.FromNode(browser.switchNode)))
	}

	if err == nil && len(screenBuffer) > 0 {
		err = ioutil.WriteFile(savepath, screenBuffer, 0666)
	}
	return err
}

//go:generate make IMPL_TYPE=FlowImplScreenshot gen-impl

func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplScreenshot{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplScreenshot) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplScreenshot) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplScreenshot) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplScreenshot) Clone() IFlowImpl {
	c := &FlowImplScreenshot{}
	_ = copier.Copy(c, impl)
	return c
}
