package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
	"github.com/naufalfmm/dayatani-farmer-api/persistents"
)

type (
	Usecases interface {
		GetByID(ctx context.Context, id uint64) (dao.Farmer, error)
		GetPaginated(ctx context.Context, req dto.FarmerListPaginationRequest) (dao.FarmerPaging, error)
		Create(ctx context.Context, req dto.CreateFarmerRequest) (dao.Farmer, error)
	}

	usecases struct {
		persists persistents.Persistents
	}
)

func Init(persist persistents.Persistents) (Usecases, error) {
	return &usecases{
		persists: persist,
	}, nil
}
