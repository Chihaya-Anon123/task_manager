package service

import (
	"strings"

	"github.com/Chihaya-Anon123/task_manager/internal/code"
	"github.com/Chihaya-Anon123/task_manager/internal/dao"
	"github.com/Chihaya-Anon123/task_manager/internal/errs"
	"github.com/Chihaya-Anon123/task_manager/internal/model"
)

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description`
	Status      string `json:"status"`
}

type CreateTaskOutput struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func CreateTask(userID uint, input CreateTaskInput) (*CreateTaskOutput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)
	input.Status = strings.TrimSpace(input.Status)

	if input.Title == "" {
		return nil, errs.New(code.CodeInvalidParams, "title is required")
	}

	if input.Status == "" {
		input.Status = "todo"
	}

	if input.Status != "todo" && input.Status != "doing" && input.Status != "done" {
		return nil, errs.New(code.CodeInvalidParams, "invalid status")
	}

	task := &model.Task{
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	}

	if err := dao.CreateTask(task); err != nil {
		return nil, errs.ErrDBError
	}

	return &CreateTaskOutput{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}, nil
}
