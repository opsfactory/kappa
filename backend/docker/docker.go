package docker

import log "github.com/Sirupsen/logrus"

type Docker struct {
}

func (dck *Docker) Poll(events <-chan string) {
	log.Info("Docker Polling")
}

func (dck *Docker) Exec(actions chan<- string) {
	log.Info("Docker Exec")
}
