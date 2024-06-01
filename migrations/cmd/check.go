package cmd

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/migrations/model"
	"github.com/naufalfmm/dayatani-farmer-api/utils/orm"
)

func checkConnection(ctx context.Context, o orm.Orm) error {
	if err := o.PingContext(ctx); err != nil {
		return err
	}

	if _, err := o.NewCreateTable().Model(&model.MigrationLog{}).IfNotExists().Exec(ctx); err != nil {
		return err
	}

	return o.PingContext(ctx)
}
