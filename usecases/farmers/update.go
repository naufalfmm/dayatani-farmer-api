package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

func (u usecases) Update(ctx context.Context, req dto.UpdateFarmerRequest) error {
	if _, err := u.persists.Repositories.Farmers.GetByID(ctx, req.ID); err != nil {
		return err
	}

	return u.persists.Repositories.Farmers.UpdateByID(ctx, req.ID, req.ToFarmer())
}
