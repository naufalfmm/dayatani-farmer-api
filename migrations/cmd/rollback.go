package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path"

	"github.com/naufalfmm/dayatani-farmer-api/migrations/model"
	"github.com/naufalfmm/dayatani-farmer-api/utils/orm"
	"github.com/urfave/cli/v2"
)

func isTargetVersionExist(ctx context.Context, o orm.Orm, targetVersion string) (bool, error) {
	if targetVersion == "" {
		return true, nil
	}

	count, err := o.
		NewSelect().
		Model(&model.MigrationLog{}).
		Where("id = ?", targetVersion).
		Count(ctx)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func rollbackVersion(ctx context.Context, o orm.Orm) (model.MigrationLog, error) {
	var (
		log model.MigrationLog
	)
	if err := o.
		NewSelect().
		Model(&log).
		Order("id DESC").
		Limit(1).
		Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return model.MigrationLog{}, nil
		}

		return model.MigrationLog{}, err
	}
	filePath := path.Join(getSQLPath(), fmt.Sprintf("%s_%s_rollback.sql", log.ID, log.Name))
	content, err := os.ReadFile(filePath)
	if err != nil {
		return model.MigrationLog{}, err
	}

	if _, err := o.
		NewRaw(string(content)).
		Exec(ctx); err != nil {
		return model.MigrationLog{}, err
	}

	if _, err := o.
		NewDelete().
		Model(&model.MigrationLog{}).
		Where("id = ?", log.ID).
		Exec(ctx); err != nil {
		return model.MigrationLog{}, err
	}

	return log, nil
}

func rollback(o orm.Orm) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if err := checkConnection(ctx.Context, o); err != nil {
			return err
		}

		ver := ctx.String("version")

		isExist, err := isTargetVersionExist(ctx.Context, o, ver)
		if err != nil {
			return err
		}

		if !isExist {
			return nil
		}

		o.Begin()
		defer o.Rollback()

		log, err := rollbackVersion(ctx.Context, o)
		if err != nil {
			return err
		}
		for log.ID != ver && log.ID != "" {
			log, err = rollbackVersion(ctx.Context, o)
			if err != nil {
				return err
			}
		}

		o.Commit()

		return nil
	}
}

func Rollback(o orm.Orm) *cli.Command {
	return &cli.Command{
		Name:    "rollback",
		Usage:   "rollback --version <version>",
		Aliases: []string{"r"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "version",
				Aliases:  []string{"v"},
				Required: false,
			},
		},
		Action: rollback(o),
	}
}
