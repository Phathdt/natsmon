package natsc

import (
	"flag"

	"github.com/nats-io/jsm.go"
	"github.com/nats-io/nats.go"
	sctx "natsmon/service-context"
)

type NatsComponent interface {
	GetManager() *jsm.Manager
}

type natsComp struct {
	id      string
	logger  sctx.Logger
	mng     *jsm.Manager
	natsURI string
}

func NewNatsComp(id string) *natsComp {
	return &natsComp{id: id}
}

func (n *natsComp) ID() string {
	return n.id
}

func (n *natsComp) InitFlags() {
	flag.StringVar(&n.natsURI, "nats-uri", "localhost:4222", "nats uri")
}

func (n *natsComp) Activate(sc sctx.ServiceContext) error {
	n.logger = sc.Logger(n.id)

	nc, err := nats.Connect(n.natsURI)
	if err != nil {
		return err
	}

	mng, err := jsm.New(nc)
	if err != nil {
		return err
	}

	n.mng = mng

	return nil
}

func (n *natsComp) Stop() error {
	return nil
}

func (n *natsComp) GetManager() *jsm.Manager {
	return n.mng
}
