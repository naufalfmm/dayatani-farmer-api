package farmers

import (
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
	"github.com/naufalfmm/dayatani-farmer-api/utils/orm"
)

type (
	Repositories interface {
	}

	repositories struct {
		orm orm.Orm
		log logger.Logger
	}
)

func Init(o orm.Orm, l logger.Logger) (Repositories, error) {
	return &repositories{
		orm: o,
		log: l,
	}, nil
}
