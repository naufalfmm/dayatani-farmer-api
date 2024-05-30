package repositories

import (
	"github.com/naufalfmm/dayatani-farmer-api/persistents/repositories/farmers"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
	"github.com/naufalfmm/dayatani-farmer-api/utils/orm"
)

type Repositories struct {
	Farmers farmers.Repositories
}

func Init(o orm.Orm, l logger.Logger) (Repositories, error) {
	f, err := farmers.Init(o, l)
	if err != nil {
		return Repositories{}, err
	}

	return Repositories{
		Farmers: f,
	}, nil
}
