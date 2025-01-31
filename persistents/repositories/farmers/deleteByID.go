package farmers

import (
	"context"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
)

func (r repositories) DeleteByID(ctx context.Context, id uint64) error {
	if _, err := r.db.GetDB().
		NewDelete().
		Model((*dao.Farmer)(nil)).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		r.log.Error(ctx, LogMsgDeleteFarmerByID).Err(err).Uint64(LogKeyID, id).Send()
		return err
	}

	return nil
}
