package container

import (
	"fmt"

	"github.com/opsfactory/kappa/container/label"
)

type ContainerReplicas []string
type ContainerBackend string

const (
	DockerBackend     ContainerBackend = "docker"
	SwarmBackend      ContainerBackend = "swarm"
	KubernetesBackend ContainerBackend = "kube"
	MesosBackend      ContainerBackend = "mesos"
)

type Container struct {
	Name            string
	Labels          label.LabelContainer
	Replicas        ContainerReplicas
	NumReplicas     int
	DesiredReplicas int
	Backend         ContainerBackend
}

func NewContainer() Container {
	return Container{}
}

func (c Container) String() string {
	return fmt.Sprintf(
		"Container{Name: %s, Labels: %s, Replicas: %s, "+
			"NumReplicas: %d, DesiredReplicas: %d, Backend: %s}",
		c.Name, c.Labels, c.Replicas, c.NumReplicas, c.DesiredReplicas, c.Backend)
}
