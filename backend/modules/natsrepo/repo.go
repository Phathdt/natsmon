package natsrepo

import (
	"context"
	"time"

	"github.com/nats-io/jsm.go"
	"github.com/nats-io/nats.go"
	"natsmon/modules/natsmodel"
)

type repo struct {
	mng *jsm.Manager
	js  nats.JetStreamContext
}

func NewRepo(mng *jsm.Manager, js nats.JetStreamContext) *repo {
	return &repo{mng: mng, js: js}
}

func (r *repo) ListJetstream(ctx context.Context, filter *jsm.StreamNamesFilter) ([]natsmodel.Jetstream, error) {
	streams, err := r.mng.Streams(filter)
	if err != nil {
		return nil, err
	}

	res := make([]natsmodel.Jetstream, len(streams))
	for i, stream := range streams {
		info, _ := stream.LatestInformation()
		res[i] = natsmodel.Jetstream{
			StreamConfig: stream.Configuration(),
			StreamInfo:   *info,
		}
	}

	return res, nil
}

func (r *repo) GetStream(ctx context.Context, stream string) (*natsmodel.Jetstream, error) {
	rs, err := r.mng.LoadStream(stream)
	if err != nil {
		return nil, err
	}

	info, err := rs.LatestInformation()
	if err != nil {
		return nil, err
	}

	return &natsmodel.Jetstream{
		StreamConfig: rs.Configuration(),
		StreamInfo:   *info,
	}, nil
}

func (r *repo) GetConsumers(ctx context.Context, stream string) ([]natsmodel.Consumer, error) {
	rs, err := r.mng.Consumers(stream)
	if err != nil {
		return nil, err
	}

	consumers := make([]natsmodel.Consumer, len(rs))
	for i, r := range rs {
		info, _ := r.LatestState()
		consumers[i] = natsmodel.Consumer{
			ConsumerConfig: r.Configuration(),
			ConsumerInfo:   info,
		}
	}

	return consumers, nil
}

func (r *repo) GetMessages(ctx context.Context, streamName string, offset int64) ([]natsmodel.Message, error) {
	consumerName := "natsmon"
	defer r.js.DeleteConsumer(streamName, consumerName)

	ackWait := 10 * time.Second
	ackPolicy := nats.AckExplicitPolicy
	maxWaiting := 1

	_, _ = r.js.AddConsumer(streamName, &nats.ConsumerConfig{
		Durable:         consumerName,
		DeliverPolicy:   nats.DeliverByStartSequencePolicy,
		OptStartSeq:     uint64(offset),
		AckPolicy:       ackPolicy,
		AckWait:         ackWait,
		MaxWaiting:      maxWaiting,
		MaxRequestBatch: 100,
	})

	sub, _ := r.js.PullSubscribe("", consumerName, nats.Bind(streamName, consumerName))

	msgs, err := sub.Fetch(100)
	if err != nil {
		return nil, err
	}

	res := make([]natsmodel.Message, len(msgs))
	for i, msg := range msgs {
		res[i] = natsmodel.ToMessage(msg)
	}

	return res, nil
}
