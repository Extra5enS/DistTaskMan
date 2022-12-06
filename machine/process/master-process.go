package state

import (
	"github.com/Extra5enS/DistTaskMan/machine/status"
)

type MasterState struct {
}

func (ms *MasterState) Run(s status.Status) status.Status {
	return 0
}
