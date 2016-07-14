package event

import (
	"fmt"

	"github.com/opsfactory/kappa/container"
)

type EventType string

const (
	ContainerStart EventType = "Start"
	ContainerDie   EventType = "Die"
)

type Event struct {
	Container container.Container
	Type      EventType
}

func (ev Event) String() string {
	return fmt.Sprintf("Event{Container: %s, Type: %s}",
		ev.Container, ev.Type)
}

func newEvent(c container.Container, t EventType) Event {
	return Event{
		Container: c,
		Type:      t,
	}
}
func NewContainerStartEvent(c container.Container) Event {
	return newEvent(c, ContainerStart)
}

func NewContainerDieEvent(c container.Container) Event {
	return newEvent(c, ContainerDie)
}
