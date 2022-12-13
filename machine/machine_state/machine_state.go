package machine_state

import (
	"log"
	"math/rand"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type MachineState struct {
	E         error
	AmIMaster bool
	Net       NetworkState
}

type NetworkState struct {
	MyLvL            int
	MyNum            uint // generates when master process starts
	UpAddress        string
	DownAddresses    []string
	MyAddress        string   `yaml:"listeningAddress"`
	UnknownAddresses []string `yaml:"networkAddresses"`
}

type yaml_sconfig struct {
	MyAddress        string   `yaml:"listeningAddress"`
	UnknownAddresses []string `yaml:"networkAddresses"`
}

func NewMasterStateFromConfig(config_file_name string) (MachineState, error) {
	m := MachineState{}
	ys := yaml_sconfig{}
	f, err := os.Open(config_file_name)
	if err != nil {
		return MachineState{}, err
	}
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&ys); err != nil {
		log.Println(config_file_name)
		return MachineState{}, err
	}
	f.Close()

	m.Net.MyAddress = ys.MyAddress
	m.Net.UnknownAddresses = ys.UnknownAddresses

	rand.Seed(time.Now().Unix())
	m.Net.MyNum = uint(rand.Int31())
	m.Net.MyLvL = 0

	return m, nil
}
