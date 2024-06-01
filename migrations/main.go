package main

import (
	"context"
	"os"

	"github.com/naufalfmm/dayatani-farmer-api/migrations/cmd"
	"github.com/naufalfmm/dayatani-farmer-api/migrations/resources/config"
	"github.com/naufalfmm/dayatani-farmer-api/migrations/resources/db"
	"github.com/naufalfmm/dayatani-farmer-api/migrations/resources/log"
	"github.com/urfave/cli/v2"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	logger, err := log.NewLogger()
	if err != nil {
		panic(err)
	}

	orm, err := db.NewPostgres(c, logger)
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Commands = []*cli.Command{
		cmd.Migrate(orm),
		cmd.Rollback(orm),
		cmd.Create(),
	}

	if err = app.Run(os.Args); err != nil {
		logger.Error(context.Background(), "when running migration").Send()
	}
}
