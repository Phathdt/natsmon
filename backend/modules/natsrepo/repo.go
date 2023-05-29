package natsrepo

import (
	"context"

	"github.com/nats-io/jsm.go"
	"natsmon/modules/natsmodel"
)

type repo struct {
	mng *jsm.Manager
}

func NewRepo(mng *jsm.Manager) *repo {
	return &repo{mng: mng}
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
