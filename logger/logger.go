package logger

import (
	"os"

	"github.com/truebluekiwi/utility/slack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
)

func Init() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.RFC3339TimeEncoder

	encoder := zapcore.NewConsoleEncoder(config)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zapcore.InfoLevel),
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(&slack.Writer{LogLevel: slack.LogLevelLog})), zapcore.ErrorLevel),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
}

func Info(args ...any) {
	logger.Info(args...)
}

func Infof(template string, args ...any) {
	logger.Infof(template, args...)
}

func Error(args ...any) {
	logger.Error(args...)
}

func Errorf(template string, args ...any) {
	logger.Errorf(template, args...)
}
