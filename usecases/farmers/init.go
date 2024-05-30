package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/persistents"
)

type (
	Usecases interface {
		GetByID(ctx context.Context, id uint64) (dao.Farmer, error)
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
