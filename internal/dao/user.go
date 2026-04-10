package dao

import (
	"errors"

	"github.com/Chihaya-Anon123/task_manager/internal/database"
	"github.com/Chihaya-Anon123/task_manager/internal/model"
	"gorm.io/gorm"
)

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("username=?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreaterUser(user *model.User) error {
	return database.DB.Create(user).Error
}
