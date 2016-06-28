package docker

import (
	"github.com/docker/engine-api/types"
	"github.com/opsfactory/kappa/container"
	"github.com/opsfactory/kappa/container/label"
	"golang.org/x/net/context"
)

func NewContainerFromDockerJSON(j types.ContainerJSON) *container.Container {
	c := container.NewContainer()
	c.Name = j.Name
	c.Labels = label.NewLabelContainerFromMap(j.Config.Labels)
	c.NumReplicas = 1
	c.DesideredReplicas = 1
	c.Replicas = container.ContainerReplicas{j.ID}
	c.Backed = container.DockerBackend
	return c
}

func (d *Docker) DockerInspect(id string) (types.ContainerJSON, error) {
	ctx, _ := context.WithCancel(context.Background())
	return d.ContainerInspect(ctx, id)
}
