package natsbiz

import (
	"context"

	"natsmon/service-context/component/natspub"
)

type publishMessageBiz struct {
	publisher natspub.NatsPubComponent
}

func NewPublishMessageBiz(publisher natspub.NatsPubComponent) *publishMessageBiz {
	return &publishMessageBiz{publisher: publisher}
}

func (b *publishMessageBiz) Response(ctx context.Context, streamName string, payload string) error {
	return b.publisher.Publish(streamName, payload)
}
