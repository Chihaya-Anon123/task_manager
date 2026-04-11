package main

import (
	"fmt"

	"github.com/Chihaya-Anon123/task_manager/internal/config"
	"github.com/Chihaya-Anon123/task_manager/internal/database"
	"github.com/Chihaya-Anon123/task_manager/internal/logger"
	"github.com/Chihaya-Anon123/task_manager/internal/router"
	"github.com/Chihaya-Anon123/task_manager/internal/service"
)

func main() {
	//读取配置
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("load config failed:%v", err))
	}

	if err := logger.InitLogger(cfg.Log); err != nil {
		panic(fmt.Sprintf("init logger failed: %v", err))
	}
	defer logger.Sync()

	if err := database.InitMySQL(cfg.Database); err != nil {
		logger.Log.Fatalw("init mysql failed", "error", err)
	}

	if err := database.AutoMigrate(); err != nil {
		logger.Log.Fatalw("auto migrate failed", "error", err)
	}

	service.InitAuthService(cfg.JWT)

	//初始化路由
	r := router.SetupRouter(cfg.JWT)

	logger.Log.Infow("server starting", "port", cfg.Server.Port)

	//启动服务
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		logger.Log.Fatalw("server run failed", "error", err)
	}
}
