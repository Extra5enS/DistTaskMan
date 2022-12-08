package process

import (
	"github.com/Extra5enS/DistTaskMan/machine/status"
)

type Process interface {
	Run(status status.Status) status.Status
}
