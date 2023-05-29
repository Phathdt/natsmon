package natsbiz

import (
	"context"

	"natsmon/modules/natsmodel"
)

type GetConsumerRepo interface {
	GetConsumers(ctx context.Context, stream string) ([]natsmodel.Consumer, error)
}

type getConsumerBiz struct {
	repo GetConsumerRepo
}

func NewGetConsumerBiz(repo GetConsumerRepo) *getConsumerBiz {
	return &getConsumerBiz{repo: repo}
}

func (b *getConsumerBiz) Response(ctx context.Context, stream string) ([]natsmodel.Consumer, error) {
	return b.repo.GetConsumers(ctx, stream)
}
