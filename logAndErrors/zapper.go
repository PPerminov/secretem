package logAndErrors

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func Sync() {
	logger.Sync()
}
func Init() {
	cfg := zap.Config{
		//Encoding: "json", // Set the encoding to JSON
		Encoding:         "console", // Set the encoding to JSON
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"stdout"}, // Output to stdout
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			// Configure the encoder as needed
			MessageKey:    "message",
			LevelKey:      "level",
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			TimeKey:       "time",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			CallerKey:     "caller",
			EncodeCaller:  zapcore.ShortCallerEncoder,
			StacktraceKey: "trace",
		},
	}

	// Create a logger instance
	logger, _ = cfg.Build()
	//logger, _ = zap.NewProduction()
	//logger.
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
