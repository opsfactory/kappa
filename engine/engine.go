package engine

import (
	log "github.com/Sirupsen/logrus"
	"github.com/opsfactory/kappa/config"
	"github.com/opsfactory/kappa/container/action"
	"github.com/opsfactory/kappa/container/backend"
	"github.com/opsfactory/kappa/container/event"
)

type Engine struct {
	cfg     config.Config
	backend backend.Backend
}

func NewEngine(cfg config.Config) (*Engine, error) {
	b, err := backend.NewBackend(cfg.Backend, cfg.BackendConfig)
	if err != nil {
		log.Fatalf("Unable to create backend %s: %v", cfg.Backend, err)
		return nil, err
	}

	return &Engine{cfg: cfg, backend: b}, nil
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

func (e Engine) handleEvent(eventsChan <-chan event.Event,
	actionsChan <-chan action.Action, errChan chan<- error) {
	for ev := range eventsChan {
		log.Infof("[EVENT] %s", ev)
	}
}
