package provider

import (
	"context"
	"github.com/sonymuhamad/todo-mailer-service/handler"
	"log"

	"github.com/sonymuhamad/todo-mailer-service/config"
	"github.com/sonymuhamad/todo-mailer-service/consumer"
)

func StartConsumer(ctx context.Context, cfg config.EnvConfig, handler *handler.BaseHandler) {
	if len(cfg.GetTopics()) == 0 {
		log.Fatal("No topic listed")

		return
	}

	for _, topic := range cfg.GetTopics() {
		kafkaConsumer := ProvideConsumer(cfg, topic, handler)

		go kafkaConsumer.Start(ctx)
	}
}

func ProvideConsumer(cfg config.EnvConfig, topic string, handler *handler.BaseHandler) *consumer.KafkaConsumer {
	return &consumer.KafkaConsumer{
		Cfg:     cfg,
		Topic:   topic,
		Handler: handler,
	}
}
