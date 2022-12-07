package state

import (
	"github.com/Extra5enS/DistTaskMan/machine/machine_state"
	"github.com/Extra5enS/DistTaskMan/machine/status"
)

type MasterState struct {
	ms *machine_state.MachineState
}

func (ms *MasterState) Run(s status.Status) status.Status {
	return 0
}

func (mms *MasterState) Init(ms *machine_state.MachineState) error {
	mms.ms = ms
	return nil
}
