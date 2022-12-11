package process

import (
	"bufio"
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
	case <-time.After(20 * time.Second):

		for _, address := range ms.ms.Net.UnknownAddresses {
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Println("Can't establish conection")
				return status.ERROR_UNKNOWEN_STATUS
			}
			fmt.Println("send")
			fmt.Fprintf(conn, `{"name":"master_req", "args":{"address":"%s", "num":"%d"}}`+"\n", ms.ms.Net.MyAddress, ms.ms.Net.MyNum)
			fmt.Println("send end")
			message, _, err := bufio.NewReader(conn).ReadLine()
			if err != nil {
				fmt.Println("Can't Read line")
				return status.ERROR_UNKNOWEN_STATUS
			}
			fmt.Printf("%s\n", message)
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
		task.Solution <- server.Sol(fmt.Sprintf(`{"my_num":"%s", "server_num":"%d"}`+"\n", task.Args[1], ms.Net.MyNum))
	},
}
