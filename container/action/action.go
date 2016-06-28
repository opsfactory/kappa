package action

import (
	"fmt"

	"github.com/opsfactory/kappa/container"
)

type ActionCommand string
type ActionUnit int

const (
	ScaleUp   ActionCommand = "ScaleUP"
	ScaleDown ActionCommand = "ScaleDown"
)

type Action struct {
	Container *container.Container
	Command   ActionCommand
	Unit      ActionUnit
}

func (a *Action) String() string {
	return fmt.Sprintf("Action{Container: %s, Command: %s, Unit: %d}",
		a.Container, a.Command, a.Unit)
}
