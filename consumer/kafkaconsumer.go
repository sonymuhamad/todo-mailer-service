package consumer

import (
	"context"
	"github.com/sonymuhamad/todo-mailer-service/handler"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/sonymuhamad/todo-mailer-service/config"
)

type KafkaConsumer struct {
	Cfg     config.EnvConfig
	Topic   string
	Handler *handler.BaseHandler
}

func (k *KafkaConsumer) Start(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: k.Cfg.GetBrokers(),
		GroupID: k.Cfg.KafkaGroupID,
		Topic:   k.Topic,
	})

	defer r.Close()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Shutting down consumer for Topic: %s", k.Topic)
			return

		default:
			m, err := r.ReadMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return // graceful shutdown
				}

				log.Printf("Error reading message from Topic %s: %v", k.Topic, err)
				continue
			}

			log.Printf("Received message from topic %s: %s", k.Topic, string(m.Value))
			errHandler := k.Handler.HandleMessage(ctx, k.Topic, m)
			if errHandler != nil {
				log.Println(errHandler, k.Topic, "")
			}
		}
	}
}
