package natsbiz

import (
	"context"

	"natsmon/modules/natsmodel"
)

type GetStreamRepo interface {
	GetStream(ctx context.Context, stream string) (*natsmodel.Jetstream, error)
}

type getStreamBiz struct {
	repo GetStreamRepo
}

func NewGetStreamBiz(repo GetStreamRepo) *getStreamBiz {
	return &getStreamBiz{repo: repo}
}

func (b *getStreamBiz) Response(ctx context.Context, stream string) (*natsmodel.Jetstream, error) {
	return b.repo.GetStream(ctx, stream)
}
