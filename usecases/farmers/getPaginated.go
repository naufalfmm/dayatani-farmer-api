package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

func (u usecases) GetPaginated(ctx context.Context, req dto.FarmerListPaginationRequest) (dao.FarmerPaging, error) {
	return u.persists.Repositories.Farmers.GetPaginated(ctx, req)
}
