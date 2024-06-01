package farmers

import (
	"testing"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/stretchr/testify/assert"
)

func Test_usecases_GetByID(t *testing.T) {
	var (
		farmer = dao.Farmer{
			ID:        1,
			Name:      "Warga Brunei Darussalam",
			BirthDate: time.Date(1977, 10, 13, 0, 0, 0, 0, time.Local),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	)

	t.Run("If get returns farmer, it will return the farmer detail", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmer.EXPECT().GetByID(mock.ctx, farmer.ID).Return(farmer, nil)

		res, err := mock.usecases.GetByID(mock.ctx, farmer.ID)

		assert.Nil(t, err)
		assert.Equal(t, farmer, res)
	})

	t.Run("If get returns error, it will return the error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmer.EXPECT().GetByID(mock.ctx, farmer.ID).Return(dao.Farmer{}, errAny)

		res, err := mock.usecases.GetByID(mock.ctx, farmer.ID)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Farmer{}, res)
	})
}
