package utilities

import "flag"

type FlagsStore struct {
	Config string
}

func FlagsStoreParse() FlagsStore {
	config := flag.String("config", "", "config file name")

	flag.Parse()
	fs := FlagsStore{}
	fs.Config = *config

	return fs
}
