package docker

import (
	"fmt"

	eventtypes "github.com/docker/engine-api/types/events"
)

type handlerFunc func(eventtypes.Message)

//// Deprecated information from JSONMessage.
//// With data only in container events.
//Status string `json:"status,omitempty"`
//ID     string `json:"id,omitempty"`
//From   string `json:"from,omitempty"`

//Type   string
//Action string
//Actor  Actor

//Time     int64 `json:"time,omitempty"`
//TimeNano int64 `json:"timeNano,omitempty"`
func createHandlerBuilder(d *Docker, ech <-chan string) handlerFunc {
	return func(m eventtypes.Message) {
		fmt.Printf("[CREATE] id: %s", m.ID)
	}
}

func destroyHandlerBuilder(d *Docker, ech <-chan string) handlerFunc {
	return func(m eventtypes.Message) {
		fmt.Printf("[DESTROY] id: %s", m.ID)
	}
}
