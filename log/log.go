package log

import (
	"go.uber.org/zap"
	"time"
	"go.uber.org/zap/zapcore"
)

func Info(level, key, value string) {
	logger.Info(level, setTime(), zap.String(key, value))
}

func Error(level, key, value string) {
	logger.Error(level, zap.String(key, value))
}

func Warn(level, key,value string) {
	logger.Warn(level, zap.String(key, value))
}

func Crit(level, key, value string) {
	logger.Fatal(level, zap.String(key, value))
}

func setTime() zapcore.Field {
	return zap.Int64("timestamp", time.Now().Unix())
}
