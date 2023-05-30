package natsbiz

import (
	"context"

	"natsmon/modules/natsmodel"
)

type GetMessageRepo interface {
	GetMessages(ctx context.Context, streamName string, offset int64) ([]natsmodel.Message, error)
}

type getMessageBiz struct {
	repo GetMessageRepo
}

func NewGetMessageBiz(repo GetMessageRepo) *getMessageBiz {
	return &getMessageBiz{repo: repo}
}

func (b *getMessageBiz) Response(ctx context.Context, streamName string, offset int64) ([]natsmodel.Message, error) {
	return b.repo.GetMessages(ctx, streamName, offset)
}
