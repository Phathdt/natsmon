package natsmodel

import "github.com/nats-io/jsm.go/api"

type Consumer struct {
	api.ConsumerConfig `json:"consumer_config"`
	api.ConsumerInfo   `json:"consumer_info"`
}
