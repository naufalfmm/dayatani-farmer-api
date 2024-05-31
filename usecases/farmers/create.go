package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

func (u usecases) Create(ctx context.Context, req dto.UpsertFarmerRequest) (dao.Farmer, error) {
	return u.persists.Repositories.Farmers.Create(ctx, req.ToFarmer())
}
