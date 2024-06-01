package log

import (
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger/zeroLogger"
)

func NewLogger() (logger.Logger, error) {
	return zeroLogger.NewZeroLogger(
		zeroLogger.WithEnabled(true),
	)
}
