package docker

import (
	"fmt"

	"golang.org/x/net/context"

	eventtypes "github.com/docker/engine-api/types/events"
	"github.com/opsfactory/kappa/backend/label"
)

type handlerFunc func(eventtypes.Message)

func startHandlerBuilder(d *Docker, ech <-chan string) handlerFunc {
	return func(m eventtypes.Message) {
		ctx, _ := context.WithCancel(context.Background())
		cj, _ := d.ContainerInspect(ctx, m.ID)

		config := cj.Config
		labels := config.Labels
		fmt.Printf("[START] Image: %s\t Name: %s\n", cj.Image, cj.Name)

		lc := label.NewLabelContainer()
		for key, label := range labels {
			fmt.Println("label: ", key, label)
		}
	}
}
