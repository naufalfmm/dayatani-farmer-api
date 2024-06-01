package farmers

import (
	"testing"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_Create(t *testing.T) {
	var (
		farmer = dao.Farmer{
			Name:      "Warga Indonesia",
			BirthDate: time.Date(2000, 04, 03, 0, 0, 0, 0, time.UTC),
		}

		id uint64 = 1
	)

	t.Run("If no error, it will return the created farmer with id", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.orm.EXPECT().NewInsert().Return(mock.ormInsert)
		mock.ormInsert.EXPECT().Model(&farmer).DoAndReturn(func(f *dao.Farmer) interface{} {
			f.ID = id
			return mock.ormInsert
		})
		mock.ormInsert.EXPECT().Returning("*").Return(mock.ormInsert)
		mock.ormInsert.EXPECT().Exec(mock.ctx).Return(nil, nil)

		expResp := farmer
		expResp.ID = id

		resp, err := mock.repositories.Create(mock.ctx, farmer)

		assert.Nil(t, err)
		assert.Equal(t, expResp, resp)
	})

	t.Run("If exec return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.orm.EXPECT().NewInsert().Return(mock.ormInsert)
		mock.ormInsert.EXPECT().Model(&farmer).Return(mock.ormInsert)
		mock.ormInsert.EXPECT().Returning("*").Return(mock.ormInsert)
		mock.ormInsert.EXPECT().Exec(mock.ctx).Return(nil, errAny)

		mock.log.EXPECT().Error(mock.ctx, LogMsgCreateFarmer).Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Any(LogKeyFarmer, farmer).Return(mock.log)
		mock.log.EXPECT().Send()

		resp, err := mock.repositories.Create(mock.ctx, farmer)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Farmer{}, resp)
	})
}
