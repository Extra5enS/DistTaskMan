package main

import (
	"math/rand"

	"github.com/Extra5enS/DistTaskMan/machine"
	"github.com/Extra5enS/DistTaskMan/machine/machine_state"
	"github.com/Extra5enS/DistTaskMan/machine/server"
)

func main() {
	tasks := make(chan server.Task)
	ms := machine_state.MachineState{
		Net: machine_state.NetworkState{
			MyNum:     uint(rand.Int31()),
			MyAddress: "0.0.0.0:3333",
		},
	}
	m, _ := machine.NewMachine(&ms, tasks)
	m.Run()
}
