package backend

import (
	"reflect"
	"testing"

	"github.com/opsfactory/kappa/config"
)

func TestNewDockerBackend(t *testing.T) {
	bc := config.BackendConfig{}
	b, err := NewBackend("docker", bc)

	if err != nil {
		t.Error(err)
	}

	bType := reflect.TypeOf(b)

	if bType.String() != "*docker.Docker" {
		t.Errorf("A backend of type *docker.Docker was expected instead of a %s", bType.String())
	}
}
