package natspub

import (
	"encoding/json"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-jetstream/pkg/jetstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/nats-io/nats.go"
	sctx "natsmon/service-context"
	"natsmon/service-context/component/common"
)

type NatsPubComponent interface {
	Publish(topic string, data interface{}) error
}

type natsPubComponent struct {
	id        string
	publisher *jetstream.Publisher
	logger    sctx.Logger
	natsURI   string
}

func NewNatsPubComponent(id string) *natsPubComponent {
	return &natsPubComponent{id: id}
}

func (n *natsPubComponent) ID() string {
	return n.id
}

func (n *natsPubComponent) InitFlags() {
	n.natsURI = common.NatsURI
}

func (n *natsPubComponent) Activate(sc sctx.ServiceContext) error {
	n.logger = sc.Logger(n.id)
	n.logger.Info("Connect to nats at ", n.natsURI, " ...")

	options := []nats.Option{
		nats.RetryOnFailedConnect(true),
		nats.Timeout(30 * time.Second),
		nats.ReconnectWait(1 * time.Second),
	}
	marshaler := &jetstream.GobMarshaler{}
	logger := watermill.NewStdLogger(false, false)

	publisher, err := jetstream.NewPublisher(
		jetstream.PublisherConfig{
			URL:         n.natsURI,
			NatsOptions: options,
			Marshaler:   marshaler,
		},
		logger,
	)
	if err != nil {
		n.logger.Error("Error connect to nats at ", n.natsURI, ". ", err.Error())
		return err
	}

	n.publisher = publisher

	return nil
}

func (n *natsPubComponent) Stop() error {
	return n.publisher.Close()
}

func (n *natsPubComponent) Publish(topic string, data interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)
	if err = n.publisher.Publish(topic, msg); err != nil {
		return err
	}

	n.logger.Infof("natsPub message = %+v\n", string(msg.Payload))

	return nil
}
