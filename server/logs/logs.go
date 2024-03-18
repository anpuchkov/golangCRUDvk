package logs

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger() *zap.Logger {
	var err error

	cfg := zap.NewProductionConfig()
	cfg.DisableCaller = true
	cfg.DisableStacktrace = true
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	logger, err = cfg.Build()

	if err != nil {
		log.Fatal("cannot init zap", err)
	}

	return logger
}

func Error(text string, fields ...zap.Field) {
	logger.Error(text, fields...)
}

func Info(text string, fields ...zap.Field) {
	logger.Info(text, fields...)
}
