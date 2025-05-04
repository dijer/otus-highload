package logger

import "go.uber.org/zap"

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

type logger struct {
	log *zap.Logger
}

func New(log *zap.Logger) Logger {
	return &logger{
		log: log,
	}
}

func (l *logger) Info(msg string, fields ...zap.Field) {
	l.log.Info(msg, fields...)
}

func (l *logger) Error(msg string, fields ...zap.Field) {
	l.log.Error(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.log.Warn(msg, fields...)
}
