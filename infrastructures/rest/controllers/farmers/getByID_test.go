package farmers

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/stretchr/testify/assert"
)

func Test_Controllers_GetByID(t *testing.T) {
	var (
		id    uint64 = 1
		idStr string = "1"

		farmer = dao.Farmer{
			ID: id,
		}
	)

	t.Run("If get farmer returns the farmer data, it will return the farmer data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetParam("id", idStr)

		mock.farmer.EXPECT().GetByID(mock.ctx, id).Return(farmer, nil)

		mock.controllers.GetByID(mock.gc)

		assert.Equal(t, http.StatusOK, mock.resRecorder.Code)
	})

	t.Run("If get farmer returns not found, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetParam("id", idStr)

		mock.farmer.EXPECT().GetByID(mock.ctx, id).Return(dao.Farmer{}, sql.ErrNoRows)

		mock.controllers.GetByID(mock.gc)

		assert.Equal(t, http.StatusBadRequest, mock.resRecorder.Code)
	})

	t.Run("If get farmer returns error unless not found, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetParam("id", idStr)

		mock.farmer.EXPECT().GetByID(mock.ctx, id).Return(dao.Farmer{}, errAny)

		mock.controllers.GetByID(mock.gc)

		assert.Equal(t, http.StatusInternalServerError, mock.resRecorder.Code)
	})

	t.Run("If id is invalid, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetParam("id", "idStr")

		mock.controllers.DeleteByID(mock.gc)

		assert.Equal(t, http.StatusBadRequest, mock.resRecorder.Code)
	})
}
