package docker

import (
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/opsfactory/kappa/container"
	"github.com/opsfactory/kappa/container/label"
	"golang.org/x/net/context"
)

func NewContainerFromDockerJSON(j types.ContainerJSON) container.Container {
	c := container.NewContainer()
	c.Name = j.Name
	c.Labels = label.NewLabelContainerFromMap(j.Config.Labels)
	c.NumReplicas = 1
	c.DesiredReplicas = 1
	c.Replicas = container.ContainerReplicas{j.ID}
	c.Backend = container.DockerBackend
	return c
}

func (d *Docker) DockerInspect(id string) (container.Container, error) {
	ctx, _ := context.WithCancel(context.Background())
	cj, err := d.ContainerInspect(ctx, id)
	if err != nil {
		log.Errorf("Error inspecting docker container %s: %v", id, err)
		return container.NewContainer(), err
	}
	return NewContainerFromDockerJSON(cj), nil
}
