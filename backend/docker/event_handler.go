package docker

import (
	log "github.com/Sirupsen/logrus"
	eventtypes "github.com/docker/engine-api/types/events"
)

type handlerFunc func(eventtypes.Message)

func startHandlerBuilder(d *Docker, ech <-chan string) handlerFunc {
	return func(m eventtypes.Message) {
		cj, err := d.DockerInspect(m.ID)
		if err != nil {
			log.Errorf("Error inspecting docker container %s: %v", m.ID, err)
		}
		c := NewContainerFromDockerJSON(cj)
		log.Printf("[START] ID: %s Name: %s Labels: %s\n", c.Replicas[0], c.Name, c.Labels)
	}
}

func dieHandlerBuilder(d *Docker, ech <-chan string) handlerFunc {
	return func(m eventtypes.Message) {
		cj, _ := d.DockerInspect(m.ID)

		log.Printf("[DIE] Image: %s Name: %s\n", cj.Image, cj.Name)
	}
}
