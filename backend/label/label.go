package label

import "reflect"

type LabelContainer struct {
	Min    string `label:"kappa.min"`
	Max    string `label:"kappa.max"`
	Rate   string `label:"kappa.rate"`
	Metric string `label:"kappa.metric"`
}

func NewLabelContainer() *LabelContainer {
	return &LabelContainer{}
}

func NewLabelContainerFromMap(m map[string]string) *LabelContainer {
	lc := NewLabelContainer()
	st := reflect.TypeOf(lc)
	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		tag := f.Tag.Get("label")
		reflect.ValueOf(f).SetString(m[tag])
	}
	return lc
}
