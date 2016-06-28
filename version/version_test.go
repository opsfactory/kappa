package version

import "testing"

func TestVersion(t *testing.T) {
	s := Version + Build + " (" + GitCommit + ")"
	v := FullVersion()

	if s != v {
		t.Errorf("'%s' does not match the expected '%s'", v, s)
	}
}
