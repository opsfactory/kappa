package backend

import (
	"fmt"
	"strings"

	"github.com/opsfactory/kappa/backend/docker"
	"github.com/opsfactory/kappa/config"
)

type Backend interface {
	Poll(events <-chan string)
	Exec(actions chan<- string)
}

func NewBackend(name string, c config.BackendConfig) (Backend, error) {

	// Invoking the Engine
	switch strings.ToLower(name) {
	case "docker":
		return &docker.Docker{}, nil
	}

	return nil, fmt.Errorf("Unknown backend %s.", name)
}
