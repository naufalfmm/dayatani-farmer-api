package frozenTime

import (
	"context"
	"sync"
	"testing"
	"time"
)

type frozenTimeKeyType string

const frozenTimeKey frozenTimeKeyType = "frozentime"

var frozenTimeMap = sync.Map{}

func Now(ctx context.Context) time.Time {
	if ctx.Value(frozenTimeKey) == nil {
		return time.Now()
	}

	t := ctx.Value(frozenTimeKey).(*testing.T)
	if t == nil {
		return time.Now()
	}

	anyNow, ok := frozenTimeMap.Load(t)
	if !ok {
		return time.Now()
	}

	now, ok := anyNow.(time.Time)
	if !ok {
		return time.Now()
	}

	return now
}

func Freeze(t *testing.T, ctx context.Context, now time.Time) context.Context {
	frozenTimeMap.Store(t, now)

	t.Cleanup(func() {
		frozenTimeMap.Delete(t)
	})

	return context.WithValue(ctx, frozenTimeKey, t)
}
