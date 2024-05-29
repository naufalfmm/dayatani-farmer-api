package controllers

import "github.com/naufalfmm/dayatani-farmer-api/usecases"

type Controllers struct{}

func Init(usc usecases.Usecases) (Controllers, error) {
	return Controllers{}, nil
}
