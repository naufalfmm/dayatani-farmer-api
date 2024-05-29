package zeroLogger

import (
	"context"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
	"github.com/rs/zerolog"
)

type zelog struct {
	zl  zerolog.Logger
	ze  *zerolog.Event
	msg string

	conf config
}

func (l *zelog) Info(ctx context.Context, msg string) logger.Logger {
	if !l.conf.enabled {
		return l
	}

	l.ze = l.zl.Info().Timestamp().Ctx(ctx)
	l.msg = msg

	return l
}

func (l *zelog) Debug(ctx context.Context, msg string) logger.Logger {
	if !l.conf.enabled {
		return l
	}

	l.ze = l.zl.Debug().Timestamp().Ctx(ctx)
	l.msg = msg

	return l
}

func (l *zelog) Warn(ctx context.Context, msg string) logger.Logger {
	if !l.conf.enabled {
		return l
	}

	l.ze = l.zl.Warn().Timestamp().Ctx(ctx)
	l.msg = msg

	return l
}

func (l *zelog) Error(ctx context.Context, msg string) logger.Logger {
	if !l.conf.enabled {
		return l
	}

	l.ze = l.zl.Error().Timestamp().Ctx(ctx).Stack()
	l.msg = msg

	return l
}

func (l *zelog) Str(key string, val string) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Str(key, val)

	return l
}

func (l *zelog) Bool(key string, val bool) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Bool(key, val)

	return l
}

func (l *zelog) Int(key string, val int) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Int(key, val)

	return l
}

func (l *zelog) Int8(key string, val int8) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Int8(key, val)

	return l
}

func (l *zelog) Int16(key string, val int16) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Int16(key, val)

	return l
}

func (l *zelog) Int32(key string, val int32) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Int32(key, val)

	return l
}

func (l *zelog) Int64(key string, val int64) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Int64(key, val)

	return l
}

func (l *zelog) Uint(key string, val uint) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Uint(key, val)

	return l
}

func (l *zelog) Uint8(key string, val uint8) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Uint8(key, val)

	return l
}

func (l *zelog) Uint16(key string, val uint16) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Uint16(key, val)

	return l
}

func (l *zelog) Uint32(key string, val uint32) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Uint32(key, val)

	return l
}

func (l *zelog) Uint64(key string, val uint64) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Uint64(key, val)

	return l
}

func (l *zelog) Float32(key string, val float32) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Float32(key, val)

	return l
}

func (l *zelog) Float64(key string, val float64) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Float64(key, val)

	return l
}

func (l *zelog) Dur(key string, val time.Duration) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Dur(key, val)

	return l
}

func (l *zelog) Time(key string, t time.Time) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Time(key, t)

	return l
}

func (l *zelog) Any(key string, val any) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Any(key, val)

	return l
}

func (l *zelog) Err(err error) logger.Logger {
	if l.ze == nil {
		return l
	}

	l.ze.Err(err)

	return l
}

func (l *zelog) Send() {
	if l.ze == nil {
		return
	}

	if l.msg != "" {
		l.ze.Send()
	} else {
		l.ze.Msg(l.msg)
	}

	l.ze = nil
	l.msg = ""
}

func (l *zelog) Printf(msg string, vs ...interface{}) {
	if l.ze == nil {
		l.Info(context.TODO(), msg)
	}

	for _, v := range vs {
		if _, ok := v.(error); ok {
			l.ze.Err(v.(error))
		}

		l.ze.Any("", v)
	}

	l.ze.Send()
}
