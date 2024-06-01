package farmers

import (
	"testing"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/utils/frozenTime"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_UpdateByID(t *testing.T) {
	var (
		id uint64 = 1
	)

	t.Run("If no error with updated data, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		updatedFarmer := dao.Farmer{
			Name:      "Warga Jerman",
			BirthDate: time.Date(1991, 12, 06, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Now(),
		}

		mock.orm.EXPECT().NewUpdate().Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Model(((*dao.Farmer)(nil))).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("name = ?", updatedFarmer.Name).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("birth_date = ?", updatedFarmer.BirthDate).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("updated_at = ?", updatedFarmer.UpdatedAt).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Where("id = ?", id).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Exec(mock.ctx).Return(nil, nil)

		err := mock.repositories.UpdateByID(mock.ctx, id, updatedFarmer)

		assert.Nil(t, err)
	})

	t.Run("If no error with missing name, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		updatedFarmer := dao.Farmer{
			BirthDate: time.Date(1991, 12, 06, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Now(),
		}

		mock.orm.EXPECT().NewUpdate().Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Model(((*dao.Farmer)(nil))).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("birth_date = ?", updatedFarmer.BirthDate).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("updated_at = ?", updatedFarmer.UpdatedAt).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Where("id = ?", id).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Exec(mock.ctx).Return(nil, nil)

		err := mock.repositories.UpdateByID(mock.ctx, id, updatedFarmer)

		assert.Nil(t, err)
	})

	t.Run("If no error with missing birth date, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		updatedFarmer := dao.Farmer{
			Name:      "Warga Jerman",
			UpdatedAt: time.Now(),
		}

		mock.orm.EXPECT().NewUpdate().Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Model(((*dao.Farmer)(nil))).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("name = ?", updatedFarmer.Name).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("updated_at = ?", updatedFarmer.UpdatedAt).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Where("id = ?", id).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Exec(mock.ctx).Return(nil, nil)

		err := mock.repositories.UpdateByID(mock.ctx, id, updatedFarmer)

		assert.Nil(t, err)
	})

	t.Run("If no error with missing updated at, it will return nil with updated at now", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		updatedFarmer := dao.Farmer{
			Name:      "Warga Jerman",
			BirthDate: time.Date(1991, 12, 06, 0, 0, 0, 0, time.UTC),
		}

		now := time.Now()
		mock.ctx = frozenTime.Freeze(t, mock.ctx, now)

		mock.orm.EXPECT().NewUpdate().Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Model(((*dao.Farmer)(nil))).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("name = ?", updatedFarmer.Name).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("birth_date = ?", updatedFarmer.BirthDate).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("updated_at = ?", now).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Where("id = ?", id).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Exec(mock.ctx).Return(nil, nil)

		err := mock.repositories.UpdateByID(mock.ctx, id, updatedFarmer)

		assert.Nil(t, err)
	})

	t.Run("If exec return error with updated data, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		updatedFarmer := dao.Farmer{
			Name:      "Warga Jerman",
			BirthDate: time.Date(1991, 12, 06, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Now(),
		}

		mock.orm.EXPECT().NewUpdate().Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Model(((*dao.Farmer)(nil))).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("name = ?", updatedFarmer.Name).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("birth_date = ?", updatedFarmer.BirthDate).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Set("updated_at = ?", updatedFarmer.UpdatedAt).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Where("id = ?", id).Return(mock.ormUpdate)
		mock.ormUpdate.EXPECT().Exec(mock.ctx).Return(nil, errAny)

		mock.log.EXPECT().Error(mock.ctx, LogMsgUpdateFarmerByID).Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Uint64(LogKeyID, id).Return(mock.log)
		mock.log.EXPECT().Send()

		err := mock.repositories.UpdateByID(mock.ctx, id, updatedFarmer)

		assert.Equal(t, errAny, err)
	})
}
