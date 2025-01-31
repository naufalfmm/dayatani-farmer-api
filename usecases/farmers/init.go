package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
	"github.com/naufalfmm/dayatani-farmer-api/persistents"
)

//go:generate mockgen -package=farmers -destination=../../mocks/usecases/farmers/init.go -source=init.go
type (
	Usecases interface {
		GetByID(ctx context.Context, id uint64) (dao.Farmer, error)
		GetPaginated(ctx context.Context, req dto.FarmerListPaginationRequest) (dao.FarmerPaging, error)
		Create(ctx context.Context, req dto.CreateFarmerRequest) (dao.Farmer, error)
		Update(ctx context.Context, req dto.UpdateFarmerRequest) error
		DeleteByID(ctx context.Context, id uint64) error
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
