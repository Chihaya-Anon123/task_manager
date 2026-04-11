package dao

import (
	"github.com/Chihaya-Anon123/task_manager/internal/database"
	"github.com/Chihaya-Anon123/task_manager/internal/model"
)

func CreateTask(task *model.Task) error {
	return database.DB.Create(task).Error
}

func ListTasksByUserID(userID uint, page, pageSize int, status string) ([]model.Task, int64, error) {
	var tasks []model.Task
	var total int64

	db := database.DB.Model(&model.Task{}).Where("user_id=?", userID)

	if status != "" {
		db = db.Where("status=?", status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := db.Order("id DESC").Offset(offset).Limit(pageSize).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}
