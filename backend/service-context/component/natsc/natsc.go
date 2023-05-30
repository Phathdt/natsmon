package natsc

import (
	"github.com/nats-io/jsm.go"
	"github.com/nats-io/nats.go"
	sctx "natsmon/service-context"
	"natsmon/service-context/component/common"
)

type NatsComponent interface {
	GetManager() *jsm.Manager
	GetJs() nats.JetStreamContext
}

type natsComp struct {
	id      string
	logger  sctx.Logger
	natsURI string
	mng     *jsm.Manager
	js      nats.JetStreamContext
}

func NewNatsComp(id string) *natsComp {
	return &natsComp{id: id}
}

func (n *natsComp) ID() string {
	return n.id
}

func (n *natsComp) InitFlags() {
	n.natsURI = common.NatsURI
}

func (n *natsComp) Activate(sc sctx.ServiceContext) error {
	n.logger = sc.Logger(n.id)

	nc, err := nats.Connect(n.natsURI)
	if err != nil {
		return err
	}

	js, err := nc.JetStream()
	if err != nil {
		return err
	}

	n.js = js

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

func (n *natsComp) GetJs() nats.JetStreamContext {
	return n.js
}
