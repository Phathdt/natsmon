package natsmodel

import "github.com/nats-io/jsm.go/api"

type Jetstream struct {
	api.StreamConfig `json:"stream_config"`
	api.StreamInfo   `json:"stream_info"`
}
