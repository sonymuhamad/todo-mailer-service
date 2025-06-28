package provider

import (
	"github.com/sonymuhamad/todo-mailer-service/config"
	"github.com/sonymuhamad/todo-mailer-service/grpcclient"
	"github.com/sonymuhamad/todo-mailer-service/handler"
	mailer2 "github.com/sonymuhamad/todo-mailer-service/mailer"
	"github.com/sonymuhamad/todo-mailer-service/usecase"
)

func ProvideHandler(cfg config.EnvConfig) *handler.BaseHandler {
	mailer := mailer2.NewMailer(cfg)
	grpcClient := grpcclient.NewGrpcClient(cfg)
	taskUsecase := usecase.NewTask(cfg, grpcClient, mailer)

	return handler.NewBaseHandler(taskUsecase)
}
