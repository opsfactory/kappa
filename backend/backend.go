package backend

import (
	"fmt"
	"strings"

	"github.com/opsfactory/kappa/backend/docker"
	"github.com/opsfactory/kappa/config"
)

type Backend interface {
	Monitor(events <-chan string)
	Exec(actions chan<- string)
}

func NewBackend(name string, c config.BackendConfig) (Backend, error) {
	var (
		b   Backend
		err error
	)

	switch strings.ToLower(name) {
	case "docker":
		b, err = docker.NewDockerBackend(c)
	default:
		err = fmt.Errorf("Unknown backend %s.", name)
	}

	if err != nil {
		return nil, err
	}
	return b, nil
}
