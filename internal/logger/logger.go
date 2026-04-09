package logger

import (
	"fmt"
	"strings"

	"github.com/Chihaya-Anon123/task_manager/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func InitLogger(cfg config.LogConfig) error {
	var zapCfg zap.Config

	//根据 format 选择开发或生产配置
	switch strings.ToLower(cfg.Format) {
	case "json":
		zapCfg = zap.NewProductionConfig()
	default:
		zapCfg = zap.NewDevelopmentConfig()
	}

	//设置日志级别
	level := new(zap.AtomicLevel)
	if err := level.UnmarshalText([]byte(strings.ToLower(cfg.Level))); err != nil {
		return fmt.Errorf("invalid log level: %w", err)
	}
	zapCfg.Level = *level

	//时间格式更直观一些
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := zapCfg.Build()
	if err != nil {
		return fmt.Errorf("build logger failed: %w", err)
	}

	Log = logger.Sugar()
	return nil
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
