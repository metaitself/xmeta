package logger

import (
	"fmt"
	"go.uber.org/zap"
)

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

func init() {
	New()
}

func SetLevel(lv string) {
	al.SetLevel(getLoggerLevel(lv))
}

func Debug(msg string, fields ...Field) {
	zl.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	zl.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	zl.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	zl.Error(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	zl.Fatal(msg, fields...)
}

func Debugf(format string, args ...any) {
	zl.Debug(fmt.Sprintf(format, args...))
}

func Infof(format string, args ...any) {
	zl.Info(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...any) {
	zl.Warn(fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...any) {
	zl.Error(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...any) {
	zl.Fatal(fmt.Sprintf(format, args...))
}

func SugaredLogger() *zap.SugaredLogger {
	return zl.Sugar()
}

type fieldLog struct {
	zl *zap.Logger
}

func WithFields(fs ...Field) Logger {
	return &fieldLog{
		zl: zl.With(fs...),
	}
}

func (l *fieldLog) Debug(msg string, fields ...Field) {
	l.zl.Debug(msg, fields...)
}

func (l *fieldLog) Info(msg string, fields ...Field) {
	l.zl.Info(msg, fields...)
}

func (l *fieldLog) Warn(msg string, fields ...Field) {
	l.zl.Warn(msg, fields...)
}

func (l *fieldLog) Error(msg string, fields ...Field) {
	l.zl.Error(msg, fields...)
}

func (l *fieldLog) Fatal(msg string, fields ...Field) {
	l.zl.Fatal(msg, fields...)
}

func (l *fieldLog) Debugf(format string, args ...any) {
	l.zl.Debug(fmt.Sprintf(format, args...))
}

func (l *fieldLog) Infof(format string, args ...any) {
	l.zl.Info(fmt.Sprintf(format, args...))
}

func (l *fieldLog) Warnf(format string, args ...any) {
	l.zl.Warn(fmt.Sprintf(format, args...))
}

func (l *fieldLog) Errorf(format string, args ...any) {
	l.zl.Error(fmt.Sprintf(format, args...))
}

func (l *fieldLog) Fatalf(format string, args ...any) {
	l.zl.Fatal(fmt.Sprintf(format, args...))
}
