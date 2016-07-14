package engine

import (
	"fmt"
	"testing"

	"github.com/opsfactory/kappa/container/action"
	"github.com/opsfactory/kappa/container/event"
)

type stubBackend struct {
}

func (b stubBackend) Monitor(eventsChan chan<- event.Event, errChan chan<- error) {
	errChan <- fmt.Errorf("I am meant to interrupt")
}

func (b stubBackend) Exec(eventsChan chan<- action.Action, errChan chan<- error) {

}

// TODO: update this tests once the actual Run function is in the process of being implemented
func TestEngineRun(t *testing.T) {
	b := stubBackend{}
	e := NewEngine(b)
	res := e.Run()

	if res == nil {
		t.Error("An error was expected")
	}
}
