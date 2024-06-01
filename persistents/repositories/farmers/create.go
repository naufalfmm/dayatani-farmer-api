package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
)

func (r repositories) Create(ctx context.Context, farmer dao.Farmer) (dao.Farmer, error) {
	if _, err := r.db.GetDB().
		NewInsert().
		Model(&farmer).
		Returning("*").
		Exec(ctx); err != nil {
		r.log.Error(ctx, LogMsgCreateFarmer).Err(err).Any(LogKeyFarmer, farmer).Send()
		return dao.Farmer{}, err
	}

	return farmer, nil
}
