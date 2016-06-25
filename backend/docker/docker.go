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
	"github.com/opsfactory/kappa/version"
	"github.com/vdemeester/docker-events"
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

func (d *Docker) Monitor(ech <-chan string) {

	ctx, _ := context.WithCancel(context.Background())

	eh := events.NewHandler(events.ByAction)

	eh.Handle("start", startHandlerBuilder(d, ech))

	filters := filters.NewArgs()
	filters.Add("type", "container")
	options := types.EventsOptions{
		Filters: filters,
	}

	errChan := events.MonitorWithHandler(ctx, d.Client, options, eh)

	if err := <-errChan; err != nil {
		// Do something
	}
}

func (d *Docker) Exec(actions chan<- string) {
	log.Info("Docker Exec")
}
