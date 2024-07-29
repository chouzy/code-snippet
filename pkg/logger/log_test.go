package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"testing"
)

var log *Logger

func setupLogger() error {
	log = NewLogger(&lumberjack.Logger{
		Filename:   "./logger.logger",
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 30,
		Compress:   false,
	}, "")

	return nil
}

func TestLog(t *testing.T) {
	err := setupLogger()
	if err != nil {
		t.Error(err)
	}
	log.Info("info...")
	log.Error("err...")
	log.Warn("warn...")
	log.Debug("debug...")
}
