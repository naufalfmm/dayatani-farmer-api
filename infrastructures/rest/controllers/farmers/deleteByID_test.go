package farmers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Controllers_DeleteByID(t *testing.T) {
	var (
		id    uint64 = 1
		idStr string = "1"
	)

	t.Run("If the deleting process is success, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetParam("id", idStr)

		mock.farmer.EXPECT().DeleteByID(mock.ctx, id).Return(nil)

		mock.controllers.DeleteByID(mock.gc)

		assert.Equal(t, http.StatusOK, mock.resRecorder.Code)
	})

	t.Run("If the deleting process returns error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetParam("id", idStr)

		mock.farmer.EXPECT().DeleteByID(mock.ctx, id).Return(errAny)

		mock.controllers.DeleteByID(mock.gc)

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
