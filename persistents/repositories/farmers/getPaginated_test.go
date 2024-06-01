package farmers

import (
	"fmt"
	"testing"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/consts"
	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_GetPaginated(t *testing.T) {
	var (
		farmers = dao.Farmers{
			{
				ID:        1,
				Name:      "Warga Indonesia",
				BirthDate: time.Date(2000, 04, 03, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:        2,
				Name:      "Warga Kosovo",
				BirthDate: time.Date(1997, 06, 23, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:        3,
				Name:      "Warga Malaysia",
				BirthDate: time.Date(1997, 06, 23, 0, 0, 0, 0, time.UTC),
			},
		}

		req = dto.FarmerListPaginationRequest{
			PaginationRequest: dto.PaginationRequest{
				Limit:  2,
				Offset: 1,
				Sorts:  []string{"-name"},
			},
		}
	)

	t.Run("If no error, it will return the paginated data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		fs := dao.Farmers{}

		mock.orm.EXPECT().NewSelect().Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Model(&fs).DoAndReturn(func(ffs *dao.Farmers) interface{} {
			*ffs = farmers[1:]
			return mock.ormSelect
		})
		mock.ormSelect.EXPECT().Order(fmt.Sprintf("%s %s", ColumnFarmerName, consts.OrdDesc)).Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Limit(req.Limit).Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Offset(req.Offset).Return(mock.ormSelect)
		mock.ormSelect.EXPECT().ScanAndCount(mock.ctx).Return(len(farmers), nil)

		expResp := dao.FarmerPaging{
			Offset: req.Offset,
			Limit:  req.Limit,
			Count:  len(farmers),
			Sorts:  req.Sorts,

			Items: farmers[1:],
		}

		resp, err := mock.repositories.GetPaginated(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, expResp, resp)
	})

	t.Run("If scan and count return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		fs := dao.Farmers{}

		mock.orm.EXPECT().NewSelect().Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Model(&fs).Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Order(fmt.Sprintf("%s %s", ColumnFarmerName, consts.OrdDesc)).Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Limit(req.Limit).Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Offset(req.Offset).Return(mock.ormSelect)
		mock.ormSelect.EXPECT().ScanAndCount(mock.ctx).Return(0, errAny)

		mock.log.EXPECT().Error(mock.ctx, LogMsgGetPaginatedFarmers).Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Any(consts.KeyReq, req).Return(mock.log)
		mock.log.EXPECT().Send()

		expResp := dao.FarmerPaging{}

		resp, err := mock.repositories.GetPaginated(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, expResp, resp)
	})
}
