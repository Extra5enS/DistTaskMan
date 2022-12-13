package main

import (
	"log"

	"github.com/Extra5enS/DistTaskMan/machine"
	"github.com/Extra5enS/DistTaskMan/machine/machine_state"
	"github.com/Extra5enS/DistTaskMan/machine/server"
	"github.com/Extra5enS/DistTaskMan/utilities"
)

func main() {
	fs := utilities.FlagsStoreParse()
	if fs.Config == "" {
		log.Fatalln("No consig")
		return
	}
	log.Println(fs)
	tasks := make(chan server.Task)
	ms, err := machine_state.NewMasterStateFromConfig(fs.Config)
	log.Println(ms)
	if err != nil {
		log.Fatalln(err)
		return
	}
	m, err := machine.NewMachine(&ms, tasks)
	if err != nil {
		log.Fatalln(err)
		return
	}
	m.Run()
}
