package controllers

import (
	"github.com/naufalfmm/dayatani-farmer-api/infrastructures/rest/controllers/farmers"
	"github.com/naufalfmm/dayatani-farmer-api/usecases"
)

type Controllers struct {
	Farmers farmers.Controllers
}

func Init(usc usecases.Usecases) (Controllers, error) {
	f, err := farmers.Init(usc)
	if err != nil {
		return Controllers{}, err
	}

	return Controllers{
		Farmers: f,
	}, nil
}
