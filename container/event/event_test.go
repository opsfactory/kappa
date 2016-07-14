package event

import (
	"testing"

	"github.com/opsfactory/kappa/container"
)

func TestNewContainerStartEvent(t *testing.T) {
	c := container.NewContainer()
	ev := NewContainerStartEvent(c)

	if ev.Type != ContainerStart {
		t.Error("event type should be `Start`, not: ", ev.Type)
	}
}

func TestNewContainerDieEvent(t *testing.T) {
	c := container.NewContainer()
	ev := NewContainerDieEvent(c)

	if ev.Type != ContainerDie {
		t.Error("event type should be `Die`, not: ", ev.Type)
	}
}
