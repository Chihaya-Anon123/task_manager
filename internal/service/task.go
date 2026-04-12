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
	Description string `json:"description"`
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

type ListTasksInput struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Status   string `form:"status"`
}

type TaskItem struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ListTasksOutput struct {
	List     []TaskItem `json:"list"`
	Total    int64      `json:"total"`
	Page     int        `json:"page"`
	PageSize int        `json:"page_size"`
}

func ListTasks(usrID uint, input ListTasksInput) (*ListTasksOutput, error) {
	if input.Page <= 0 {
		input.Page = 1
	}

	if input.PageSize <= 0 {
		input.PageSize = 10
	}

	if input.PageSize > 100 {
		input.PageSize = 100
	}

	input.Status = strings.TrimSpace(input.Status)
	if input.Status != "" &&
		input.Status != "todo" &&
		input.Status != "doing" &&
		input.Status != "done" {
		return nil, errs.New(code.CodeInvalidParams, "invalid status")
	}

	tasks, total, err := dao.ListTasksByUserID(usrID, input.Page, input.PageSize, input.Status)
	if err != nil {
		return nil, errs.ErrDBError
	}

	list := make([]TaskItem, 0, len(tasks))
	for _, task := range tasks {
		list = append(list, TaskItem{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
		})
	}

	return &ListTasksOutput{
		List:     list,
		Total:    total,
		Page:     input.Page,
		PageSize: input.PageSize,
	}, nil
}

type TaskDetailOutput struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func GetTaskDetail(taskID, userID uint) (*TaskDetailOutput, error) {
	if taskID == 0 {
		return nil, errs.New(code.CodeInvalidParams, "invalid task id")
	}

	task, err := dao.GetTaskByIDAndUserID(taskID, userID)
	if err != nil {
		return nil, errs.ErrDBError
	}
	if task == nil {
		return nil, errs.New(code.CodeNotFound, "task not found")
	}

	return &TaskDetailOutput{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}, nil
}

type UpdateTaskInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

type UpdateTaskOutput struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func UpdateTask(userID, taskID uint, input UpdateTaskInput) (*UpdateTaskOutput, error) {
	if taskID == 0 {
		return nil, errs.New(code.CodeInvalidParams, "invalid task id")
	}

	task, err := dao.GetTaskByIDAndUserID(taskID, userID)
	if err != nil {
		return nil, errs.ErrDBError
	}
	if task == nil {
		return nil, errs.New(code.CodeNotFound, "task not found")
	}

	if input.Title != nil {
		title := strings.TrimSpace(*input.Title)
		if title == "" {
			return nil, errs.New(code.CodeInvalidParams, "title cannot be empty")
		}
		task.Title = title
	}

	if input.Description != nil {
		task.Description = strings.TrimSpace(*input.Description)
	}

	if input.Status != nil {
		status := strings.TrimSpace(*input.Status)
		if status != "todo" && status != "doing" && status != "done" {
			return nil, errs.New(code.CodeInvalidParams, "invalid status")
		}
		task.Status = status
	}

	if err := dao.UpdateTask(task); err != nil {
		return nil, errs.ErrDBError
	}

	return &UpdateTaskOutput{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}, nil
}

func DeleteTask(userID, taskID uint) error {
	if taskID == 0 {
		return errs.New(code.CodeInvalidParams, "invalid task id")
	}

	task, err := dao.GetTaskByIDAndUserID(taskID, userID)
	if err != nil {
		return errs.ErrDBError
	}
	if task == nil {
		return errs.New(code.CodeNotFound, "task not found")
	}

	if err := dao.DeleteTask(task); err != nil {
		return errs.ErrDBError
	}

	return nil
}
