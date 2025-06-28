package grpcclient

import (
	"context"
	"time"

	taskpb "github.com/sonymuhamad/todo-mailer-service/protogen/task"
)

type TaskClient struct {
	*BaseGrpcClient
	Client taskpb.TaskServiceClient
}

func NewTaskClient(target string) (*TaskClient, error) {
	baseClient, err := NewBaseGrpcClient(target)
	if err != nil {
		return nil, err
	}

	return &TaskClient{
		BaseGrpcClient: baseClient,
		Client:         taskpb.NewTaskServiceClient(baseClient.Conn),
	}, nil
}

func (c *TaskClient) Close() error { return c.BaseGrpcClient.Conn.Close() }

func (c *TaskClient) GetTaskByID(id string) (*taskpb.TaskResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return c.Client.GetTaskByID(ctx, &taskpb.GetTaskByIDRequest{Id: id})
}
