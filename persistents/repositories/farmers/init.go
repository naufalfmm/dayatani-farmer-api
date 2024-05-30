package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/resources/db"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
)

type (
	Repositories interface {
		GetByID(ctx context.Context, id uint64) (dao.Farmer, error)
	}

	repositories struct {
		db  *db.DB
		log logger.Logger
	}
)

func Init(d *db.DB, l logger.Logger) (Repositories, error) {
	return &repositories{
		db:  d,
		log: l,
	}, nil
}
