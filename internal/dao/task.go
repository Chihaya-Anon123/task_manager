package dao

import (
	"github.com/Chihaya-Anon123/task_manager/internal/database"
	"github.com/Chihaya-Anon123/task_manager/internal/model"
)

func CreateTask(task *model.Task) error {
	return database.DB.Create(task).Error
}
