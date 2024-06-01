package farmers

import (
	"testing"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_DeleteByID(t *testing.T) {
	var (
		id uint64 = 1
	)

	t.Run("If no error, it will return no error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.orm.EXPECT().NewDelete().Return(mock.ormDelete)
		mock.ormDelete.EXPECT().Model((*dao.Farmer)(nil)).Return(mock.ormDelete)
		mock.ormDelete.EXPECT().Where("id = ?", id).Return(mock.ormDelete)
		mock.ormDelete.EXPECT().Exec(mock.ctx).Return(nil, nil)

		err := mock.repositories.DeleteByID(mock.ctx, id)

		assert.Nil(t, err)
	})

	t.Run("If scan return error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.orm.EXPECT().NewDelete().Return(mock.ormDelete)
		mock.ormDelete.EXPECT().Model((*dao.Farmer)(nil)).Return(mock.ormDelete)
		mock.ormDelete.EXPECT().Where("id = ?", id).Return(mock.ormDelete)
		mock.ormDelete.EXPECT().Exec(mock.ctx).Return(nil, errAny)

		mock.log.EXPECT().Error(mock.ctx, LogMsgDeleteFarmerByID).Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Uint64(LogKeyID, id).Return(mock.log)
		mock.log.EXPECT().Send()

		err := mock.repositories.DeleteByID(mock.ctx, id)

		assert.Equal(t, errAny, err)
	})
}
