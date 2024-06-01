package db

import (
	"github.com/naufalfmm/dayatani-farmer-api/migrations/resources/config"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
	"github.com/naufalfmm/dayatani-farmer-api/utils/orm"
	"github.com/naufalfmm/dayatani-farmer-api/utils/orm/driver/pgOrm"
)

func NewPostgres(c *config.EnvConfig, log logger.Logger) (orm.Orm, error) {
	return pgOrm.NewPostgres(
		pgOrm.WithHostPort(c.DbHost, c.DbPort),
		pgOrm.WithUsernamePassword(c.DbUsername, c.DbPassword),
		pgOrm.WithDatabaseName(c.DbName),
		pgOrm.WithRetry(c.DbRetry, c.DbWaitSleep),
		pgOrm.WithSSLMode(c.DbSslMode),
		pgOrm.WithLogger(log),
	)
}
