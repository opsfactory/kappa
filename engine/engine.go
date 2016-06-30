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

var (
	eventChan  chan (event.Event)
	actionChan chan (action.Action)
)

func NewEngine(cfg config.Config) (*Engine, error) {

	b, err := backend.NewBackend(cfg.Backend, cfg.BackendConfig)
	if err != nil {
		log.Fatalf("Unable to create backend %s: %v", cfg.Backend, err)
		return nil, err
	}

	return &Engine{cfg: cfg, backend: b}, nil

}

func (e *Engine) Run() {

	// channel setup
	eventChan := make(chan event.Event)
	actionChan := make(chan action.Action)

	go e.backend.Monitor(eventChan)
	go e.backend.Exec(actionChan)
	for ev := range eventChan {
		log.Infof("[EVENT] %s", ev)
	}

}
