package db

import (
	"github.com/naufalfmm/dayatani-farmer-api/resources/config"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
	"github.com/naufalfmm/dayatani-farmer-api/utils/orm/driver/pgOrm"
)

func NewPostgres(c *config.EnvConfig, log logger.Logger) (*DB, error) {
	confs := []pgOrm.PostgresConfig{
		pgOrm.WithHostPort(c.DbHost, c.DbPort),
		pgOrm.WithUsernamePassword(c.DbUsername, c.DbPassword),
		pgOrm.WithDatabaseName(c.DbName),
		pgOrm.WithSessionTimeout(c.DbSessionTimeout),
		pgOrm.WithStatementTimeout(c.DbStatementTimeout),
		pgOrm.WithTransactionSessionTimeout(c.DbTransactionSessionTimeout),
		pgOrm.WithRetry(c.DbRetry, c.DbWaitSleep),
		pgOrm.WithSSLMode(c.DbSslMode),
	}

	if c.DbLogMode {
		confs = append(confs, pgOrm.WithLog(log, c.DbLogSlowThreshold))
	}

	o, err := pgOrm.NewPostgres(confs...)
	if err != nil {
		return nil, err
	}

	return &DB{
		o: o,
	}, nil
}
