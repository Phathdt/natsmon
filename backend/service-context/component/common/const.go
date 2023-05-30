package common

import "github.com/namsral/flag"

var (
	NatsURI = ""
)

func init() {
	flag.StringVar(&NatsURI, "nats-uri", "localhost:4222", "nats uri")
	flag.Parse()
}
