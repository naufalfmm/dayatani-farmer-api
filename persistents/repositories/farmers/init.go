package farmers

import (
	"github.com/naufalfmm/dayatani-farmer-api/resources/db"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
)

type (
	Repositories interface {
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
