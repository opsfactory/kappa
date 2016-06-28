package docker

import (
	log "github.com/Sirupsen/logrus"
	eventtypes "github.com/docker/engine-api/types/events"
	"github.com/opsfactory/kappa/backend/label"
	"golang.org/x/net/context"
)

type handlerFunc func(eventtypes.Message)

func startHandlerBuilder(d *Docker, ech <-chan string) handlerFunc {
	return func(m eventtypes.Message) {
		ctx, _ := context.WithCancel(context.Background())
		cj, _ := d.ContainerInspect(ctx, m.ID)

		lc := label.NewLabelContainerFromMap(cj.Config.Labels)
		log.Printf("[START] Image: %s Name: %s Labels: %s\n", cj.Image, cj.Name, lc)
	}
}

func dieHandlerBuilder(d *Docker, ech <-chan string) handlerFunc {
	return func(m eventtypes.Message) {
		ctx, _ := context.WithCancel(context.Background())
		cj, _ := d.ContainerInspect(ctx, m.ID)

		log.Printf("[DIE] Image: %s Name: %s\n", cj.Image, cj.Name)
	}
}
