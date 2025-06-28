package usecase

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/segmentio/kafka-go"
	"github.com/sonymuhamad/todo-mailer-service/config"
	"github.com/sonymuhamad/todo-mailer-service/dto"
	"github.com/sonymuhamad/todo-mailer-service/grpcclient"
	"github.com/sonymuhamad/todo-mailer-service/mailer"
)

type Task struct {
	mailer     *mailer.Mailer
	grpcClient *grpcclient.GrpcClient
	cfg        config.EnvConfig
}

func NewTask(cfg config.EnvConfig,
	client *grpcclient.GrpcClient,
	mailer *mailer.Mailer,
) *Task {
	return &Task{
		mailer:     mailer,
		grpcClient: client,
		cfg:        cfg,
	}
}

type Message struct {
	TaskID string `json:"task_id"`
}

func (t *Task) HandleMessage(ctx context.Context, message kafka.Message) error {
	var m Message

	if err := json.Unmarshal(message.Value, &m); err != nil {
		return err
	}

	resTask, err := t.grpcClient.TaskClient.GetTaskByID(m.TaskID)
	if err != nil {
		return err
	}

	if resTask == nil {
		return errors.New("task is not found")
	}

	resUser, err := t.grpcClient.UserClient.GetUserByID(resTask.UserId)
	if err != nil {
		return err
	}

	if resUser == nil {
		return errors.New("user is not found")
	}

	todos := make([]dto.CreateTodoNotificationParam, 0)
	for _, td := range resTask.Todos {
		todos = append(todos, dto.CreateTodoNotificationParam{
			Name: td.Name,
		})
	}

	err = t.mailer.TaskCreatedNotification(ctx, dto.CreateTaskNotificationParam{
		MailerHeaderParam: dto.MailerHeaderParam{
			To:      []string{resUser.Email},
			From:    []string{t.cfg.SMTPFrom},
			Subject: "Create task notification",
		},
		TaskName:        resTask.Name,
		TaskDescription: resTask.Description,
		Todos:           todos,
	})
	if err != nil {
		return err
	}

	return nil
}
