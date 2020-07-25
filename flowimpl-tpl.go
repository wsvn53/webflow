// +build ignore

package main
import "reflect"
import "strings"
import "github.com/jinzhu/copier"
func init() {
	flowImpl := func() IFlowImpl {
		return &FlowImplBase{}
	}()
	registerFlow(flowImpl)
}

func (impl *FlowImplBase) Name() string {
	interfaceName := reflect.TypeOf(impl).String()
	commandName := strings.Split(interfaceName, "FlowImpl")[1]
	return strings.ToLower(commandName)
}

func (impl *FlowImplBase) SetCommand(command *FlowCommand) {
	impl.command = command
}

func (impl *FlowImplBase) Command() *FlowCommand {
	return impl.command
}

func (impl *FlowImplBase) Clone() IFlowImpl {
	c := &FlowImplBase{}
	_ = copier.Copy(c, impl)
	return c
}
