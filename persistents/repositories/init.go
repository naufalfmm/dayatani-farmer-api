package repositories

import (
	"github.com/naufalfmm/dayatani-farmer-api/persistents/repositories/farmers"
	"github.com/naufalfmm/dayatani-farmer-api/resources/db"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
)

type Repositories struct {
	Farmers farmers.Repositories
}

func Init(d *db.DB, l logger.Logger) (Repositories, error) {
	f, err := farmers.Init(d, l)
	if err != nil {
		return Repositories{}, err
	}

	return Repositories{
		Farmers: f,
	}, nil
}
