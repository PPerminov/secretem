package logAndErrors

import (
	"go.uber.org/zap"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func Sync() {
	logger.Sync()
}
func Init() {
	logger, _ = zap.NewProduction()
	sugar = logger.Sugar()
}

func Info(msg string) {
	sugar.Infow(msg)
}
func Warn(msg string) {
	sugar.Warn(msg)
}
func Error(msg string) {
	sugar.Error(msg)
}
func Fatal(msg string) {
	sugar.Fatal(msg)
}
