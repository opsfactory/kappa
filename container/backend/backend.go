package backend

import (
	"fmt"
	"strings"

	"github.com/opsfactory/kappa/config"
	"github.com/opsfactory/kappa/container"
	"github.com/opsfactory/kappa/container/action"
	"github.com/opsfactory/kappa/container/backend/docker"
	"github.com/opsfactory/kappa/container/event"
)

type Backend interface {
	Monitor(events chan<- event.Event)
	Exec(actions chan<- action.Action)
}

func NewBackend(name string, c config.BackendConfig) (Backend, error) {
	switch strings.ToLower(name) {
	case container.DockerBackend:
		return docker.NewDockerBackend(c)
	}
	return nil, fmt.Errorf("Unknown backend %s.", name)
}
