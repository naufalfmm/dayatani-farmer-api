package farmers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_usecases_DeleteByID(t *testing.T) {
	var (
		id uint64 = 1
	)

	t.Run("If delete farmer is success, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmer.EXPECT().DeleteByID(mock.ctx, id).Return(nil)

		err := mock.usecases.DeleteByID(mock.ctx, id)

		assert.Nil(t, err)
	})

	t.Run("If delete farmer returns error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmer.EXPECT().DeleteByID(mock.ctx, id).Return(errAny)

		err := mock.usecases.DeleteByID(mock.ctx, id)

		assert.Equal(t, errAny, err)
	})
}
