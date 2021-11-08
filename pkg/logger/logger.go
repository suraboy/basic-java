package logger

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLog    *zap.Logger
	AppLogger *Logger
)

type Logger struct {
	logger *zap.SugaredLogger
}

func Init() {
	initZap()
	AppLogger = &Logger{logger: zapLog.Sugar()}
}

func initZap() {
	var err error
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLog, err = loggerConfig.Build()
	if err != nil {
		log.Fatalln("Got error while building zap config.")
		return
	}
	return
}

func (l *Logger) Request(IP, method, path, body, query string) {
	l.logger.Infow(fmt.Sprintf("request method %s path %s", method, path),
		"IP", IP,
		"method", method,
		"path", path,
		"body", body,
		"query", query,
	)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Error(v ...interface{}) {
	l.logger.Error(v)
}
