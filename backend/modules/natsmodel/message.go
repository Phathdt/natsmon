package natsmodel

import (
	"strings"

	"github.com/ThreeDotsLabs/watermill-jetstream/pkg/jetstream"
	"github.com/nats-io/nats.go"
)

type Message struct {
	Seq     string `json:"seq"`
	UUID    string `json:"uuid"`
	Payload string `json:"payload"`
}

func ToMessage(m *nats.Msg) Message {
	seq := strings.Split(m.Reply, ".")[5]
	marshaler := &jetstream.GobMarshaler{}

	message, _ := marshaler.Unmarshal(m)

	return Message{
		Seq:     seq,
		UUID:    message.UUID,
		Payload: string(message.Payload),
	}
}
