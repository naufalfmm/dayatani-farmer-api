package logger

import (
	"context"
	"time"
)

//go:generate mockgen -package=mockLogger -destination=./mockLogger/mock.go -source=logger.go
type Logger interface {
	Info(ctx context.Context, msg string) Logger
	Debug(ctx context.Context, msg string) Logger
	Warn(ctx context.Context, msg string) Logger
	Error(ctx context.Context, msg string) Logger

	Str(key string, val string) Logger
	Bool(key string, val bool) Logger
	Int(key string, val int) Logger
	Int8(key string, val int8) Logger
	Int16(key string, val int16) Logger
	Int32(key string, val int32) Logger
	Int64(key string, val int64) Logger
	Uint(key string, val uint) Logger
	Uint8(key string, val uint8) Logger
	Uint16(key string, val uint16) Logger
	Uint32(key string, val uint32) Logger
	Uint64(key string, val uint64) Logger
	Float32(key string, val float32) Logger
	Float64(key string, val float64) Logger
	Dur(key string, val time.Duration) Logger
	Time(key string, t time.Time) Logger
	Any(key string, val any) Logger
	Err(err error) Logger

	Send()

	Printf(string, ...interface{})
}
