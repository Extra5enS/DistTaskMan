package machine_state

type MachineState struct {
	E         error
	AmIMaster bool
	Net       NetworkState
}

type NetworkState struct {
	MyLvL         int
	MyNum         uint // generates when master process starts
	MyAddress     string
	UpAddress     string
	DownAddresses []string
}
