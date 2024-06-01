package farmers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
	"github.com/stretchr/testify/assert"
)

func Test_Controllers_GetPaginated(t *testing.T) {
	var (
		req = dto.FarmerListPaginationRequest{
			PaginationRequest: dto.PaginationRequest{
				Limit:  1,
				Offset: 2,
				Sorts:  []string{"name"},
			},
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetURL(fmt.Sprintf("?limit=%d&offset=%d&sorts=%s", req.Limit, req.Offset, req.Sorts[0]))

		mock.farmer.EXPECT().GetPaginated(mock.ctx, req).Return(dao.FarmerPaging{}, nil)

		mock.controllers.GetPaginated(mock.gc)

		assert.Equal(t, http.StatusOK, mock.resRecorder.Code)
	})

	t.Run("If get paginated returns error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.SetURL(fmt.Sprintf("?limit=%d&offset=%d&sorts=%s", req.Limit, req.Offset, req.Sorts[0]))

		mock.farmer.EXPECT().GetPaginated(mock.ctx, req).Return(dao.FarmerPaging{}, errAny)

		mock.controllers.GetPaginated(mock.gc)

		assert.Equal(t, http.StatusInternalServerError, mock.resRecorder.Code)
	})
}
