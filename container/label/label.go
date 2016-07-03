package label

import (
	"fmt"
	"reflect"
)

type LabelContainer struct {
	Min    string `label:"kappa.min"`
	Max    string `label:"kappa.max"`
	Rate   string `label:"kappa.rate"`
	Metric string `label:"kappa.metric"`
}

func (lc LabelContainer) String() string {
	return fmt.Sprintf(
		"LabelContainer{Min: %s, Max: %s, Rate: %s, Metric: %s}",
		lc.Min, lc.Max, lc.Rate, lc.Metric)
}

func NewLabelContainer() LabelContainer {
	return LabelContainer{}
}

func NewLabelContainerFromMap(m map[string]string) LabelContainer {

	lc := NewLabelContainer()
	st := reflect.TypeOf(lc)
	sv := reflect.ValueOf(&lc).Elem()
	for i := 0; i < st.NumField(); i++ {
		ft := st.Field(i)
		fv := sv.Field(i)
		tag := ft.Tag.Get("label")
		fv.SetString(m[tag])
	}
	return lc
}
