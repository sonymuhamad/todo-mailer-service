package dto

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Handler interface {
	HandleMessage(ctx context.Context, message kafka.Message) error
}
