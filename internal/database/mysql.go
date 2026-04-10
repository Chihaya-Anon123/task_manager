package database

import (
	"fmt"

	"github.com/Chihaya-Anon123/task_manager/internal/config"
	"github.com/Chihaya-Anon123/task_manager/internal/logger"
	"github.com/Chihaya-Anon123/task_manager/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL(cfg config.DatabaseConfig) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connect database failed: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get generic database object failed: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("ping database failed: %w", err)
	}

	DB = db
	logger.Log.Infow("mysql connected successfully",
		"host", cfg.Host,
		"port", cfg.Port,
		"dbname", cfg.DBName,
	)
	return nil
}

func AutoMigrate() error {
	if DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	if err := DB.AutoMigrate(&model.Task{}, &model.User{}); err != nil {
		return fmt.Errorf("auto migrate failed: %w", err)
	}

	logger.Log.Infow("database mmigrate successfully",
		"table", "tasks",
	)
	return nil
}
