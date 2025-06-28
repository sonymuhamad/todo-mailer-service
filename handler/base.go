package handler

import (
	"context"
	"errors"

	"github.com/segmentio/kafka-go"
	"github.com/sonymuhamad/todo-mailer-service/usecase"
)

type BaseHandler struct {
	taskUsecase *usecase.Task
}

func NewBaseHandler(taskUsecase *usecase.Task) *BaseHandler {
	return &BaseHandler{
		taskUsecase: taskUsecase,
	}
}

func (b *BaseHandler) HandleMessage(ctx context.Context, topic string, message kafka.Message) error {
	switch topic {
	case "send-mail":
		return b.taskUsecase.HandleMessage(ctx, message)
	}

	return errors.New("Unhandled topic")
}
