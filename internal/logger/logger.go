package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var l *zap.SugaredLogger
var once sync.Once

func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		l = initLogger()
	})
	return l
}

func initLogger() *zap.SugaredLogger {
	opts := make([]zap.Option, 0)

	//opts = append(opts, zap.AddCaller())
	cfg := zap.NewDevelopmentConfig()
	level := zap.DebugLevel

	cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(
				cfg.EncoderConfig,
			),
			os.Stderr,
			level,
		),
		opts...,
	).Sugar()
}
