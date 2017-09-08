package log

import (
	"go.uber.org/zap"
	"time"
	"go.uber.org/zap/zapcore"
)

const key = "content"

// TODO(fk): logger should be used more convenient

func Info(level, value string) {
	logger.Info(level, setTime(), zap.String(key, value))
}

func Error(level, value string) {
	logger.Error(level, zap.String(key, value))
}

func Warn(level, value string) {
	logger.Warn(level, zap.String(key, value))
}

func Crit(level, value string) {
	logger.Fatal(level, zap.String(key, value))
}

func setTime() zapcore.Field {
	return zap.Int64("timestamp", time.Now().Unix())
}
