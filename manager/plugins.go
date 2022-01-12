package manager

import "time"

type Manager interface {
	Start() error
	Stop() error
}

func RunManager() Manager {
	// TODO: Load interval from config.
	return &heimdallPlugin{CheckInterval: time.Second * 5}
}
