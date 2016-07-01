// Package docker implements the docker-engine backend.
package docker

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/go-connections/sockets"
	"github.com/docker/go-connections/tlsconfig"
	"github.com/opsfactory/kappa/config"
	kaction "github.com/opsfactory/kappa/container/action"
	kevent "github.com/opsfactory/kappa/container/event"
	"github.com/vdemeester/docker-events"

	"github.com/opsfactory/kappa/version"
)

const (
	DockerAPI        = "unix:///var/run/docker.sock"
	DockerAPIVersion = "v1.23"
)

type Docker struct {
	*client.Client
}

func NewDockerBackend(c config.BackendConfig) (*Docker, error) {
	var httpClient *http.Client
	httpHeaders := map[string]string{
		"User-Agent": fmt.Sprintf("kappa/%s", version.Version),
	}
	if c.TLSCert != "" && c.TLSKey != "" {
		tlsOptions := tlsconfig.Options{
			CAFile:             c.TLSCACert,
			CertFile:           c.TLSCert,
			KeyFile:            c.TLSKey,
			InsecureSkipVerify: c.AllowInsecure,
		}
		config, err := tlsconfig.Client(tlsOptions)
		if err != nil {
			return nil, err
		}
		tr := &http.Transport{
			TLSClientConfig: config,
		}
		proto, addr, _, err := client.ParseHost(DockerAPI)
		if err != nil {
			return nil, err
		}
		sockets.ConfigureTransport(tr, proto, addr)
		httpClient = &http.Client{
			Transport: tr,
		}
	}

	client, err := client.NewClient(DockerAPI, DockerAPIVersion, httpClient, httpHeaders)
	if err != nil {
		return nil, err
	}

	return &Docker{client}, nil
}

func (d *Docker) Monitor(eventsChan chan<- kevent.Event, errChan chan<- error) {

	log.Debug("[Docker][Monitor] Start")

	eh := events.NewHandler(events.ByAction)
	eh.Handle("start", startHandlerBuilder(d, eventsChan))
	eh.Handle("die", dieHandlerBuilder(d, eventsChan))

	filters := filters.NewArgs()
	filters.Add("type", "container")
	options := types.EventsOptions{
		Filters: filters,
	}

	ctx, _ := context.WithCancel(context.Background())
	hErrChan := events.MonitorWithHandler(ctx, d.Client, options, eh)
	for err := range hErrChan {
		if err != nil {
			log.Errorf("Error polling Docker events: %v.", err)
		}
	}
}

func (d *Docker) Exec(actionsChan chan<- kaction.Action, errChan chan<- error) {
	log.Debug("[Docker][Exec] Start")
}
