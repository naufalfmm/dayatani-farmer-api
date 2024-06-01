package farmers

import (
	"testing"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
	"github.com/stretchr/testify/assert"
)

func Test_usecases_GetPaginated(t *testing.T) {
	var (
		req = dto.FarmerListPaginationRequest{
			PaginationRequest: dto.PaginationRequest{
				Limit:  1,
				Offset: 1,
				Sorts:  []string{"-name"},
			},
		}
	)

	t.Run("If get paginated returns the paging farmers, it will return the paging farmers", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		farmerPaging := dao.FarmerPaging{
			Limit:  req.Limit,
			Offset: req.Offset,
			Count:  3,
			Sorts:  req.Sorts,
			Items: dao.Farmers{
				{
					ID:        2,
					Name:      "Warga Vietnam",
					BirthDate: time.Date(1966, 11, 05, 0, 0, 0, 0, time.UTC),
				},
			},
		}

		mock.farmer.EXPECT().GetPaginated(mock.ctx, req).Return(farmerPaging, nil)

		res, err := mock.usecases.GetPaginated(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, farmerPaging, res)
	})

	t.Run("If get paginated returns error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmer.EXPECT().GetPaginated(mock.ctx, req).Return(dao.FarmerPaging{}, errAny)

		res, err := mock.usecases.GetPaginated(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.FarmerPaging{}, res)
	})
}
