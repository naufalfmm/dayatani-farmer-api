package farmers

import (
	"context"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
)

func (r repositories) UpdateByID(ctx context.Context, id uint64, updatedFarmer dao.Farmer) error {
	updOrm := r.db.GetDB().
		NewUpdate().
		Model((*dao.Farmer)(nil))

	if updatedFarmer.Name != "" {
		updOrm = updOrm.Set("name = ?", updatedFarmer.Name)
	}

	if !updatedFarmer.BirthDate.IsZero() {
		updOrm = updOrm.Set("birth_date = ?", updatedFarmer.BirthDate)
	}

	updatedAt := time.Now()
	if !updatedFarmer.UpdatedAt.IsZero() {
		updatedAt = updatedFarmer.UpdatedAt
	}

	if _, err := updOrm.
		Set("updated_at = ?", updatedAt).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		r.log.Error(ctx, "update-farmer-by-id").Err(err).Uint64("id", id).Send()
		return err
	}

	return nil
}
