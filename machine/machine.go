package machine

import (
	"github.com/Extra5enS/DistTaskMan/machine/machine_state"
	"github.com/Extra5enS/DistTaskMan/machine/process"
	"github.com/Extra5enS/DistTaskMan/machine/server"
	"github.com/Extra5enS/DistTaskMan/machine/status"
)

type Machine struct {
	ms    *machine_state.MachineState
	s     server.Server
	tasks chan server.Task
}

func (m *Machine) Run() error {
	end := make(chan error)
	m.s.Run(end)
	pr := process.NewMasterProcess(m.ms, m.tasks)
	for {
		pr.Run(status.OK_STATUS)
	}
	return nil
}

func NewMachine(ms *machine_state.MachineState, tasks chan server.Task) (Machine, error) {
	m := Machine{}
	m.ms = ms
	m.tasks = tasks
	s, err := server.NewServer(m.ms.Net.MyAddress, m.tasks)
	if err != nil {
		return Machine{}, err
	}
	m.s = s
	return m, nil
}
