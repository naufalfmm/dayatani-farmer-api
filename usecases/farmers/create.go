package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
)

func (u usecases) Create(ctx context.Context, farmer dao.Farmer) (dao.Farmer, error) {
	return u.persists.Repositories.Farmers.Create(ctx, farmer)
}
