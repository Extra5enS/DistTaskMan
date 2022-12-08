package process

import (
	"github.com/Extra5enS/DistTaskMan/machine/machine_state"
	"github.com/Extra5enS/DistTaskMan/machine/server"
	"github.com/Extra5enS/DistTaskMan/machine/status"
)

type masterProcess struct {
	ms    *machine_state.MachineState
	tasks chan server.Task
}

func (ms *masterProcess) Run(s status.Status) status.Status {
	return 0
}

func NewMasterProcess(ms *machine_state.MachineState, tasks chan server.Task) masterProcess {
	mms := masterProcess{}
	mms.ms = ms
	mms.tasks = tasks
	return mms
}
