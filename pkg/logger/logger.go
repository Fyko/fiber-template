package logger

import "go.uber.org/zap"

var log *zap.SugaredLogger

func CreateLogger() *zap.SugaredLogger {
	if log != nil {
		return log
	}

	logger, _ := zap.NewDevelopment()
	log = logger.Sugar()

	return log
}
