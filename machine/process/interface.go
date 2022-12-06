package state

import (
	"github.com/Extra5enS/DistTaskMan/machine/status"
)

type State interface {
	Run(status status.Status) status.Status
}
