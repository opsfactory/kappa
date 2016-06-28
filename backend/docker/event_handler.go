package docker

import (
	log "github.com/Sirupsen/logrus"
	eventtypes "github.com/docker/engine-api/types/events"
	kappaevent "github.com/opsfactory/kappa/container/event"
)

type handlerFunc func(eventtypes.Message)

func startHandlerBuilder(d *Docker, ech chan<- *kappaevent.Event) handlerFunc {
	return func(m eventtypes.Message) {
		c, err := d.DockerInspect(m.ID)
		if err != nil {
			return
		}
		log.Infof("[START] ID: %s Name: %s Labels: %s\n", c.Replicas[0], c.Name, c.Labels)
		ech <- kappaevent.NewContainerStartEvent(c)
	}
}

func dieHandlerBuilder(d *Docker, ech chan<- *kappaevent.Event) handlerFunc {
	return func(m eventtypes.Message) {
		c, err := d.DockerInspect(m.ID)
		if err != nil {
			return
		}
		log.Printf("[DIE] ID: %s Name: %s\n", c.Replicas[0], c.Name)
		ech <- kappaevent.NewContainerDieEvent(c)
	}
}
