package main

import (
	"log"

	"github.com/opsfactory/kappa/config"
)

func main() {
	data := `
backend: docker
backend_config:
  TLSCACert: "/etc/docker/ca.crt"
  TLSCert: "/etc/docker/server.pem"
  TLSKey: "/etc/docker/server.key"
  AllowInsecure: false
metrics:
  queue_length: "/usr/local/bin/queue_length.sh"
  reqsec:  "/usr/local/bin/reqsec.sh"
`

	c, err := config.Parse(data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config.Print(c)
}
