package dto

type MailerHeaderParam struct {
	To      []string
	From    []string
	Subject string
}

type CreateTaskNotificationParam struct {
	MailerHeaderParam

	TaskName        string
	TaskDescription string
	Todos           []CreateTodoNotificationParam
}

type CreateTodoNotificationParam struct {
	Name string
}
