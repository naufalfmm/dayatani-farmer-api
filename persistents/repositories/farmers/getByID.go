package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
)

func (r repositories) GetByID(ctx context.Context, id uint64) (dao.Farmer, error) {
	var farmer dao.Farmer
	if err := r.db.GetDB().
		NewSelect().
		Model(&farmer).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		r.log.Error(ctx, LogMsgGetFarmerByID).Err(err).Uint64(LogKeyID, id).Send()
		return dao.Farmer{}, err
	}

	return farmer, nil
}
