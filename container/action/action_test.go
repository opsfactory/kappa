package action

import (
	"testing"

	"github.com/opsfactory/kappa/container"
	"github.com/opsfactory/kappa/container/label"
)

func TestActionString(t *testing.T) {
	ct := container.NewContainer()
	ct.Name = "john"
	ct.Backend = container.MesosBackend
	ct.Replicas = []string{}
	ct.NumReplicas = 0
	ct.DesiredReplicas = 2
	lbls := label.NewLabelContainer()
	lbls.Max = "10"
	lbls.Min = "1"
	lbls.Rate = "1"
	lbls.Metric = "mymetric"
	ct.Labels = lbls

	command := ScaleUp
	var unit ActionUnit
	unit = 1
	a := Action{
		Container: &ct,
		Command:   command,
		Unit:      unit,
	}

	estr := "Action{Container: Container{Name: john, Labels: LabelContainer{Min: 1, Max: 10, Rate: 1, Metric: mymetric}, Replicas: [], NumReplicas: 0, DesiredReplicas: 2, Backend: mesos}, Command: ScaleUP, Unit: 1}"

	res := a.String()
	if res != estr {
		t.Errorf(
			"Expected string and Action string does not match:\nExpected:%s\nGiven:%s",
			estr,
			res,
		)
	}
}
