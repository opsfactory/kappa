package engine

import (
	log "github.com/Sirupsen/logrus"
	"github.com/opsfactory/kappa/container/action"
	"github.com/opsfactory/kappa/container/backend"
	"github.com/opsfactory/kappa/container/event"
)

type Engine struct {
	backend backend.Backend
}

func NewEngine(b backend.Backend) *Engine {
	return &Engine{backend: b}
}

func (e Engine) Run() error {
	// channel setup
	errChan := make(chan error)
	eventsChan := make(chan event.Event)
	actionsChan := make(chan action.Action)

	go e.backend.Monitor(eventsChan, errChan)
	go e.backend.Exec(actionsChan, errChan)
	go e.handleEvent(eventsChan, actionsChan, errChan)

	for err := range errChan {
		log.Errorf("Unexpected error: %v.", err)
		return err
	}
	return nil
}

func (e Engine) handleEvent(
	eventsChan <-chan event.Event,
	actionsChan <-chan action.Action,
	errChan chan<- error,
) {
	for ev := range eventsChan {
		log.Infof("[EVENT] %s", ev)
	}
}
