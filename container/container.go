package container

import "github.com/opsfactory/kappa/container/label"

type ContainerReplicas []string
type ContainerBackend string

const (
	DockerBackend     ContainerBackend = "docker"
	SwarmBackend      ContainerBackend = "swarm"
	KubernetesBackend ContainerBackend = "kube"
	MesosBackedn      ContainerBackend = "mesos"
)

type Container struct {
	Name              string
	Labels            *label.LabelContainer
	Replicas          ContainerReplicas
	NumReplicas       int
	DesideredReplicas int
	Backed            ContainerBackend
}

func NewContainer() *Container {
	return &Container{}
}
