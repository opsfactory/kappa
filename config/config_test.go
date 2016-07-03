package config

import "testing"

func TestParse(t *testing.T) {

	yml := []byte(`
backend: docker
backend_config:
  TLSCACert: "/etc/docker/ca.crt"
  TLSCert: "/etc/docker/server.pem"
  TLSKey: "/etc/docker/server.key"
  AllowInsecure: false
metrics:
  queue_length: "/usr/local/bin/queue_length.sh"
  reqsec:  "/usr/local/bin/reqsec.sh"
`)
	c, err := parse(yml)

	if err != nil {
		t.Error(err)
	}

	if c.Backend != "docker" {
		t.Error("`backend` was supposed to be `docker`")
	}

	if c.BackendConfig.AllowInsecure != false {
		t.Error("`backend_config.AllowInsecure` is supposed to be false")
	}

	expectedTLSKey := "/etc/docker/server.key"
	if c.BackendConfig.TLSKey != expectedTLSKey {
		t.Errorf("`backend_config.TLSKey` is supposed to be %s", expectedTLSKey)
	}

	expectedTLSCert := "/etc/docker/server.pem"
	if c.BackendConfig.TLSCert != expectedTLSCert {
		t.Errorf("`backend_config.TLSCert ` is supposed to be %s", expectedTLSCert)
	}

	expectedTLSCACert := "/etc/docker/ca.crt"
	if c.BackendConfig.TLSCACert != expectedTLSCACert {
		t.Errorf("`backend_config.TLSCACert` is supposed to be %s", expectedTLSCACert)
	}

	if _, ok := c.Metrics["queue_length"]; !ok {
		t.Error("A metric named `queue_length` was expected")
	}

	if c.Metrics["queue_length"] != "/usr/local/bin/queue_length.sh" {
		t.Error("Given queue_length values is wrong")
	}

	if _, ok := c.Metrics["reqsec"]; !ok {
		t.Error("A metric named `reqsec` was expected")
	}

	if c.Metrics["reqsec"] != "/usr/local/bin/reqsec.sh" {
		t.Error("Given reqsec values is wrong")
	}
}
