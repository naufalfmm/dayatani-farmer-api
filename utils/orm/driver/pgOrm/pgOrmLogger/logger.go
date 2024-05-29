package pgOrmLogger

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger/zeroLogger"
	"github.com/uptrace/bun"
)

type LogQueryHook struct {
	enabled       bool
	logger        logger.Logger
	slowThreshold time.Duration
}

func NewLogQueryHook(opts ...Option) *LogQueryHook {
	defLogger, _ := zeroLogger.NewZeroLogger()
	h := &LogQueryHook{
		enabled: true,
		logger:  defLogger,
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func (h *LogQueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *LogQueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	if !h.enabled {
		return
	}

	if h.logger == nil {
		return
	}

	dur := time.Since(event.StartTime)

	if h.slowThreshold > 0 && dur > h.slowThreshold {
		h.logger.Warn(ctx, fmt.Sprintf("query run slower than %d", h.slowThreshold)).Str("query", event.Query).Dur("duration", dur).Send()
	}

	switch event.Err {
	case nil, sql.ErrNoRows, sql.ErrTxDone:
		h.logger.Info(ctx, "").Str("query", event.Query).Dur("duration", dur).Send()
	default:
		h.logger.Error(ctx, "").Err(event.Err).Str("query", event.Query).Dur("duration", dur).Send()
	}
}

type Option func(*LogQueryHook)

func WithLogger(logger logger.Logger) Option {
	return func(h *LogQueryHook) {
		h.logger = logger
	}
}

func WithSlowThreshold(slowThreshold time.Duration) Option {
	return func(h *LogQueryHook) {
		h.slowThreshold = slowThreshold
	}
}
