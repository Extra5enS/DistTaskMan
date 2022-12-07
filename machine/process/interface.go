package state

import (
	"github.com/Extra5enS/DistTaskMan/machine/machine_state"
	"github.com/Extra5enS/DistTaskMan/machine/status"
)

type State interface {
	Run(status status.Status) status.Status
	Init(ms *machine_state.MachineState) error
}
