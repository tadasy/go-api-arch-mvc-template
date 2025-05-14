package logger

import (
	"os"

	"go.uber.org/zap"
)

var (
	ZapLogger       *zap.Logger
	zapSugardLogger *zap.SugaredLogger
)

func init() {
	cfg := zap.NewProductionConfig()
	logFile := os.Getenv("APP_LOG_FILE")
	if logFile != "" {
		cfg.OutputPaths = []string{"stderr", logFile}
	}

	ZapLogger = zap.Must(cfg.Build())
	if os.Getenv("APP_ENV") == "development" {
		ZapLogger = zap.Must(zap.NewDevelopment())
	}
	zapSugardLogger = ZapLogger.Sugar()
}

func Sync() {
	if err := zapSugardLogger.Sync(); err != nil {
		zap.Error(err)
	}
}

func Debug(msg string, keysAndValues ...interface{}) {
	zapSugardLogger.Debugw(msg, keysAndValues...)
}
func Info(msg string, keysAndValues ...interface{}) {
	zapSugardLogger.Infow(msg, keysAndValues...)
}
func Warn(msg string, keysAndValues ...interface{}) {
	zapSugardLogger.Warnw(msg, keysAndValues...)
}
func Error(msg string, keysAndValues ...interface{}) {
	zapSugardLogger.Errorw(msg, keysAndValues...)
}
func Fatal(msg string, keysAndValues ...interface{}) {
	zapSugardLogger.Fatalw(msg, keysAndValues...)
}
func Panic(msg string, keysAndValues ...interface{}) {
	zapSugardLogger.Panicw(msg, keysAndValues...)
}
