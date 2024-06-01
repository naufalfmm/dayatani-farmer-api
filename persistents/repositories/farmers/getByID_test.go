package farmers

import (
	"testing"
	"time"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_GetByID(t *testing.T) {
	var (
		farmer = dao.Farmer{
			ID:        1,
			Name:      "Warga Indonesia",
			BirthDate: time.Date(2000, 04, 03, 0, 0, 0, 0, time.UTC),
		}
	)

	t.Run("If no error, it will return the farmer data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var frm dao.Farmer
		mock.orm.EXPECT().NewSelect().Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Model(&frm).DoAndReturn(func(f *dao.Farmer) interface{} {
			*f = farmer
			return mock.ormSelect
		})
		mock.ormSelect.EXPECT().Where("id = ?", farmer.ID).Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Scan(mock.ctx).Return(nil)

		resp, err := mock.repositories.GetByID(mock.ctx, farmer.ID)

		assert.Nil(t, err)
		assert.Equal(t, farmer, resp)
	})

	t.Run("If scan return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var frm dao.Farmer
		mock.orm.EXPECT().NewSelect().Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Model(&frm).Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Where("id = ?", farmer.ID).Return(mock.ormSelect)
		mock.ormSelect.EXPECT().Scan(mock.ctx).Return(errAny)

		mock.log.EXPECT().Error(mock.ctx, LogMsgGetFarmerByID).Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Uint64(LogKeyID, farmer.ID).Return(mock.log)
		mock.log.EXPECT().Send()

		resp, err := mock.repositories.GetByID(mock.ctx, farmer.ID)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Farmer{}, resp)
	})
}
