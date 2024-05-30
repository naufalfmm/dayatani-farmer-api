package usecases

import (
	"github.com/naufalfmm/dayatani-farmer-api/persistents"
	"github.com/naufalfmm/dayatani-farmer-api/usecases/farmers"
)

type Usecases struct {
	Farmers farmers.Usecases
}

func Init(persists persistents.Persistents) (Usecases, error) {
	f, err := farmers.Init(persists)
	if err != nil {
		return Usecases{}, err
	}

	return Usecases{
		Farmers: f,
	}, nil
}
