package main

import (
	"context"
	"github.com/sonymuhamad/todo-mailer-service/provider"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sonymuhamad/todo-mailer-service/config"
)

func main() {
	cfg := config.LoadEnvConfig()

	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	handler := provider.ProvideHandler(cfg)
	log.Printf("Starting mailer service:")

	go func() {
		provider.StartConsumer(ctx, cfg, handler)
	}()

	sig := <-sigChan
	log.Printf("Received shutdown signal: %s", sig)

	// Trigger cancellation
	cancel()

	log.Println("Shutdown Gracefully")
}
