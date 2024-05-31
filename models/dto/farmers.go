package dto

import (
	"github.com/naufalfmm/dayatani-farmer-api/consts"
	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
)

type FarmerResponse struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
}

func NewFarmerResponse(f dao.Farmer) FarmerResponse {
	return FarmerResponse{
		ID:        f.ID,
		Name:      f.Name,
		BirthDate: f.BirthDate.Format(consts.LayoutISOTime),
	}
}
