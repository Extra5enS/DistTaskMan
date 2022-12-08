package process

import (
	"fmt"
	"net"
	"time"

	"github.com/Extra5enS/DistTaskMan/machine/machine_state"
	"github.com/Extra5enS/DistTaskMan/machine/server"
	"github.com/Extra5enS/DistTaskMan/machine/status"
)

type masterProcess struct {
	ms    *machine_state.MachineState
	tasks chan server.Task
}

func (ms *masterProcess) Run(s status.Status) status.Status {
	select {
	case t := <-ms.tasks:
		go MasterTasks[t.Req](t, ms.ms)
	case <-time.After(5 * time.Second):
		for _, address := range ms.ms.Net.UnknownAddresses {
			conn, _ := net.Dial("tcp", address)
		}
	}
	return status.OK_STATUS
}

func NewMasterProcess(ms *machine_state.MachineState, tasks chan server.Task) masterProcess {
	mms := masterProcess{}
	mms.ms = ms
	mms.tasks = tasks
	return mms
}

var MasterTasks = map[string]func(task server.Task, ms *machine_state.MachineState){
	"master_req": func(task server.Task, ms *machine_state.MachineState) {
		fmt.Println(task)
		task.Solution <- server.Sol(fmt.Sprintf("Sent num:%s, Server num:%d", task.Args[1], ms.Net.MyNum))
	},
}
