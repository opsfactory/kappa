package backend

import (
	"fmt"
	"strings"

	"github.com/opsfactory/kappa/backend/docker"
	"github.com/opsfactory/kappa/config"
	kappaevent "github.com/opsfactory/kappa/container/event"
)

type Backend interface {
	Monitor(events chan<- kappaevent.Event)
	Exec(actions chan<- string)
}

func NewBackend(name string, c config.BackendConfig) (Backend, error) {
	switch strings.ToLower(name) {
	case "docker":
		return docker.NewDockerBackend(c)
	}

	return nil, fmt.Errorf("Unknown backend %s.", name)
}
