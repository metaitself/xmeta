package logger

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

// Option custom setup config
type Option func(*option)

type option struct {
	level          zapcore.Level
	file           io.Writer
	encode         string
	timeLayout     string
	callerSkip     int
	addCaller      bool
	disableConsole bool
}

// WithLevel only greater than 'level' will output
func WithLevel(level string) Option {
	return func(opt *option) {
		opt.level = getLoggerLevel(level)
	}
}

func WithCaller(enabled bool) Option {
	return func(opt *option) {
		opt.addCaller = enabled
	}
}

func AddCallerSkip(skip int) Option {
	return func(opt *option) {
		opt.callerSkip = skip
	}
}

// WithFileP write log to some file
func WithFileP(file string) Option {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0766); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0766)
	if err != nil {
		panic(err)
	}

	return func(opt *option) {
		opt.file = zapcore.Lock(f)
	}
}

// WithFileRotationP write log to some file with rotation
func WithFileRotationP(file string) Option {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0766); err != nil {
		panic(err)
	}

	return func(opt *option) {
		opt.file = &lumberjack.Logger{
			Filename:   file, // 文件路径
			MaxSize:    128,  // 单个文件最大尺寸，默认单位 M
			MaxBackups: 300,  // 最多保留 300 个备份
			MaxAge:     30,   // 最大时间，默认单位 day
			LocalTime:  true, // 使用本地时间
			Compress:   true, // 是否压缩 disabled by default
		}
	}
}

// WithTimeLayout custom time format
func WithTimeLayout(timeLayout string) Option {
	return func(opt *option) {
		opt.timeLayout = timeLayout
	}
}

// WithDisableConsole WithEnableConsole write log to os.Stdout or os.Stderr
func WithDisableConsole() Option {
	return func(opt *option) {
		opt.disableConsole = true
	}
}

func WithEncode(encode string) Option {
	return func(opt *option) {
		opt.encode = encode
	}
}
