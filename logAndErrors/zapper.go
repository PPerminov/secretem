package logAndErrors

import (
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func Init() {
	logger = log.New()

}

func Info(msg string) {
	logger.Info(msg)
}
func Warn(msg string) {
	logger.Warn(msg)
}
func Error(msg string) {
	logger.Error(msg)
}
func Fatal(msg string) {
	logger.Fatal(msg)
	//logger.Fatal(logger.Trace())
}
