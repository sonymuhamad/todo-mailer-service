package grpcclient

import (
	"github.com/sonymuhamad/todo-mailer-service/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type BaseGrpcClient struct {
	Conn *grpc.ClientConn
}

func NewBaseGrpcClient(target string) (*BaseGrpcClient, error) {
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		return nil, err
	}

	return &BaseGrpcClient{
		Conn: conn,
	}, nil
}

type GrpcClient struct {
	UserClient *UserClient
	TaskClient *TaskClient
}

func NewGrpcClient(cfg config.EnvConfig) *GrpcClient {
	taskClient, err := NewTaskClient(cfg.TaskServiceGrpcServer)
	if err != nil {
		log.Fatal(err)

		return nil
	}

	userClient, err := NewUserClient(cfg.UserServiceGrpcServer)
	if err != nil {
		log.Fatal(err)

		return nil
	}

	return &GrpcClient{
		UserClient: userClient,
		TaskClient: taskClient,
	}
}
