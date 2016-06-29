package container

import (
	"fmt"

	"github.com/opsfactory/kappa/container/label"
)

const (
	DockerBackend     string = "docker"
	SwarmBackend      string = "swarm"
	KubernetesBackend string = "kube"
	MesosBackend      string = "mesos"
)

type ContainerReplicas []string

type Container struct {
	Name            string
	Labels          label.LabelContainer
	Replicas        ContainerReplicas
	NumReplicas     int
	DesiredReplicas int
	Backend         string
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
