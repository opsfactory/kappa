package docker

import (
	log "github.com/Sirupsen/logrus"
	eventtypes "github.com/docker/engine-api/types/events"
	kappaevent "github.com/opsfactory/kappa/container/event"
)

type handlerFunc func(eventtypes.Message)

func startHandlerBuilder(d *Docker, ech chan<- kappaevent.Event) handlerFunc {
	return func(m eventtypes.Message) {
		c, err := d.Inspect(m.ID)
		if err != nil {
			return
		}
		log.Debugf("[Docker][EventHandler][Start] ID: %s Name: %s Labels: %s", c.Replicas[0], c.Name, c.Labels)
		ech <- kappaevent.NewContainerStartEvent(&c)
	}
}

func dieHandlerBuilder(d *Docker, ech chan<- kappaevent.Event) handlerFunc {
	return func(m eventtypes.Message) {
		c, err := d.Inspect(m.ID)
		if err != nil {
			return
		}
		log.Debugf("[Docker][EventHandler][Die] ID: %s Name: %s", c.Replicas[0], c.Name)
		ech <- kappaevent.NewContainerDieEvent(&c)
	}
}
