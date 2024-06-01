package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
	"github.com/naufalfmm/dayatani-farmer-api/resources/db"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
)

const (
	LogMsgCreateFarmer        = "create-farmer"
	LogMsgGetFarmerByID       = "get-farmer-by-id"
	LogMsgDeleteFarmerByID    = "delete-farmer-by-id"
	LogMsgGetPaginatedFarmers = "get-paginated-farmers"
	LogMsgUpdateFarmerByID    = "update-farmer-by-id"

	LogKeyFarmer = "farmer"
	LogKeyID     = "id"

	ColumnFarmerName = "farmer.name"
)

//go:generate mockgen -package=farmers -destination=../../../mocks/persistents/repositories/farmers/init.go -source=init.go
type (
	Repositories interface {
		GetByID(ctx context.Context, id uint64) (dao.Farmer, error)
		GetPaginated(ctx context.Context, req dto.FarmerListPaginationRequest) (dao.FarmerPaging, error)
		Create(ctx context.Context, farmer dao.Farmer) (dao.Farmer, error)
		UpdateByID(ctx context.Context, id uint64, updatedFarmer dao.Farmer) error
		DeleteByID(ctx context.Context, id uint64) error
	}

	repositories struct {
		db  *db.DB
		log logger.Logger
	}
)

func Init(d *db.DB, l logger.Logger) (Repositories, error) {
	return &repositories{
		db:  d,
		log: l,
	}, nil
}
