package farmers

import (
	"testing"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
	"github.com/naufalfmm/dayatani-farmer-api/utils/frozenTime"
	"github.com/stretchr/testify/assert"
)

func Test_usecases_Update(t *testing.T) {
	var (
		now = time.Now()

		req = dto.UpdateFarmerRequest{
			ID: 1,
			CreateFarmerRequest: dto.CreateFarmerRequest{
				Name:      "Warga Timor Leste",
				BirthDate: time.Date(1988, 10, 05, 0, 0, 0, 0, time.UTC),
			},
		}
	)

	t.Run("If the updating farmer is success, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.ctx = frozenTime.Freeze(t, mock.ctx, now)

		existingFarmer := dao.Farmer{
			ID:        req.ID,
			Name:      "Warga Singapura",
			BirthDate: time.Date(1987, 10, 05, 0, 0, 0, 0, time.UTC),
			CreatedAt: now.Add(-1 * time.Hour),
			UpdatedAt: now.Add(-1 * time.Hour),
		}

		updatedFarmer := dao.Farmer{
			Name:      req.Name,
			BirthDate: req.BirthDate,
			CreatedAt: now,
			UpdatedAt: now,
		}

		mock.farmer.EXPECT().GetByID(mock.ctx, req.ID).Return(existingFarmer, nil)
		mock.farmer.EXPECT().UpdateByID(mock.ctx, req.ID, updatedFarmer).Return(nil)

		err := mock.usecases.Update(mock.ctx, req)

		assert.Nil(t, err)
	})

	t.Run("If the updating farmer is error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.ctx = frozenTime.Freeze(t, mock.ctx, now)

		existingFarmer := dao.Farmer{
			ID:        req.ID,
			Name:      "Warga Singapura",
			BirthDate: time.Date(1987, 10, 05, 0, 0, 0, 0, time.UTC),
			CreatedAt: now.Add(-1 * time.Hour),
			UpdatedAt: now.Add(-1 * time.Hour),
		}

		updatedFarmer := dao.Farmer{
			Name:      req.Name,
			BirthDate: req.BirthDate,
			CreatedAt: now,
			UpdatedAt: now,
		}

		mock.farmer.EXPECT().GetByID(mock.ctx, req.ID).Return(existingFarmer, nil)
		mock.farmer.EXPECT().UpdateByID(mock.ctx, req.ID, updatedFarmer).Return(errAny)

		err := mock.usecases.Update(mock.ctx, req)

		assert.Equal(t, errAny, err)
	})

	t.Run("If get farmer error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmer.EXPECT().GetByID(mock.ctx, req.ID).Return(dao.Farmer{}, errAny)

		err := mock.usecases.Update(mock.ctx, req)

		assert.Equal(t, errAny, err)
	})
}
