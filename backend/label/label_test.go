package label

import "testing"

func TestNewLabelContainerFromMap(t *testing.T) {
	m := map[string]string{
		"kappa.min":    "10",
		"kappa.max":    "40",
		"kappa.rate":   "4",
		"kappa.metric": "queue_length",
	}

	lc := NewLabelContainerFromMap(m)
	givenString := lc.String()
	expectedString := "LabelContainer{Min: 10, Max: 40, Rate: 4, Metric: queue_length}"
	if givenString != expectedString {
		t.Errorf("'%s' does not match the expected '%s'", givenString, expectedString)
	}

	if lc.Min != "10" {
		t.Errorf("lc.Min is %s which is not equal to 10", lc.Min)
	}

	if lc.Max != "40" {
		t.Errorf("lc.Max is %s which is not equal to 40", lc.Max)
	}

	if lc.Rate != "4" {
		t.Errorf("lc.Rate is %s which is not equal to 4", lc.Rate)
	}

	if lc.Metric != "queue_length" {
		t.Errorf("lc.Metric is %s which is not equal to 'queue_length'", lc.Metric)
	}
}
