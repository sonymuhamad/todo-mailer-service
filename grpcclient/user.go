package grpcclient

import (
	"context"
	"time"

	userpb "github.com/sonymuhamad/todo-mailer-service/protogen/user"
)

type UserClient struct {
	*BaseGrpcClient
	Client userpb.UserServiceClient
}

func NewUserClient(target string) (*UserClient, error) {
	baseClient, err := NewBaseGrpcClient(target)
	if err != nil {
		return nil, err
	}

	return &UserClient{
		BaseGrpcClient: baseClient,
		Client:         userpb.NewUserServiceClient(baseClient.Conn),
	}, nil
}

func (c *UserClient) Close() error { return c.BaseGrpcClient.Conn.Close() }

func (c *UserClient) GetUserByID(id string) (*userpb.GetUserByIDResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return c.Client.GetUserByID(ctx, &userpb.GetUserByIDRequest{Id: id})
}
