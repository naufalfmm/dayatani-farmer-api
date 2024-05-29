package pgOrm

import (
	"database/sql"

	"github.com/naufalfmm/dayatani-farmer-api/utils/orm"
	"github.com/naufalfmm/dayatani-farmer-api/utils/orm/driver/pgOrm/pgOrmLogger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewPostgres(configs ...PostgresConfig) (orm.Orm, error) {
	conf, err := generateDefault()
	if err != nil {
		return nil, err
	}

	for _, config := range configs {
		config(&conf)
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(conf.toDriverOptions()...))
	db := bun.NewDB(sqldb, pgdialect.New())

	for i := 0; i < conf.retry; i++ {
		if err := db.Ping(); err == nil {
			break
		}
	}

	if conf.logger != nil {
		db.AddQueryHook(pgOrmLogger.NewLogQueryHook(
			pgOrmLogger.WithLogger(conf.logger),
			pgOrmLogger.WithSlowThreshold(conf.logSlowThreshold),
		))
	}

	return &postgresOrm{
		db: db,
	}, nil
}
