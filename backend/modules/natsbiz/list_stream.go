package natsbiz

import (
	"context"

	"github.com/nats-io/jsm.go"
	"natsmon/modules/natsmodel"
)

type listJetstreamRepo interface {
	ListJetstream(ctx context.Context, filter *jsm.StreamNamesFilter) ([]natsmodel.Jetstream, error)
}

type listJetstreamBiz struct {
	repo listJetstreamRepo
}

func NewListJetstreamBiz(repo listJetstreamRepo) *listJetstreamBiz {
	return &listJetstreamBiz{repo: repo}
}

func (b *listJetstreamBiz) Response(ctx context.Context, filter *jsm.StreamNamesFilter) ([]natsmodel.Jetstream, error) {
	return b.repo.ListJetstream(ctx, filter)
}
