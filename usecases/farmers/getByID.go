package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
)

func (u usecases) GetByID(ctx context.Context, id uint64) (dao.Farmer, error) {
	return u.persists.Repositories.Farmers.GetByID(ctx, id)
}
