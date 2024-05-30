package log

import (
	"github.com/naufalfmm/dayatani-farmer-api/resources/config"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger/zeroLogger"
)

func NewLogger(c *config.EnvConfig) (logger.Logger, error) {
	return zeroLogger.NewZeroLogger(
		zeroLogger.WithEnabled(c.LogMode),
	)
}
