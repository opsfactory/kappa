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
	Container *container.Container
	Type      EventType
}

func (ev Event) String() string {
	return fmt.Sprintf("Event{Container: %s, Type: %s}",
		ev.Container, ev.Type)
}

func NewEvent() *Event {
	return &Event{}
}

func NewContainerStartEvent(c *container.Container) *Event {
	e := NewEvent()
	e.Container = c
	e.Type = ContainerStart
	return e
}

func NewContainerDieEvent(c *container.Container) *Event {
	e := NewEvent()
	e.Container = c
	e.Type = ContainerDie
	return e
}
