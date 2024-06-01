package farmers

import (
	"testing"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
	"github.com/naufalfmm/dayatani-farmer-api/utils/frozenTime"
	"github.com/stretchr/testify/assert"
)

func Test_usecases_Create(t *testing.T) {
	var (
		now = time.Now()

		req = dto.CreateFarmerRequest{
			Name:      "Warga Tajikistan",
			BirthDate: time.Date(1970, 11, 21, 0, 0, 0, 0, time.UTC),
		}
	)

	t.Run("If the creating process is success, it will return the created farmer", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		fr := dao.Farmer{
			Name:      req.Name,
			BirthDate: req.BirthDate,
			CreatedAt: now,
			UpdatedAt: now,
		}

		returnedFarmer := fr
		returnedFarmer.ID = 1

		mock.ctx = frozenTime.Freeze(t, mock.ctx, now)

		mock.farmer.EXPECT().Create(mock.ctx, fr).Return(returnedFarmer, nil)

		resp, err := mock.usecases.Create(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, returnedFarmer, resp)
	})

	t.Run("If the creating process returns error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		fr := dao.Farmer{
			Name:      req.Name,
			BirthDate: req.BirthDate,
			CreatedAt: now,
			UpdatedAt: now,
		}

		mock.ctx = frozenTime.Freeze(t, mock.ctx, now)

		mock.farmer.EXPECT().Create(mock.ctx, fr).Return(dao.Farmer{}, errAny)

		resp, err := mock.usecases.Create(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Farmer{}, resp)
	})
}
