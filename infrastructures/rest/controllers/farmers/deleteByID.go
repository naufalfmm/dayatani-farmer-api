package farmers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

func (c Controllers) DeleteByID(gc *gin.Context) {
	id, err := strconv.ParseUint(gc.Param("id"), 10, 64)
	if err != nil {
		gc.JSON(http.StatusBadRequest, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data: dto.ErrorData{
				Error: err.Error(),
			},
		})

		return
	}

	if err := c.Usecases.Farmers.DeleteByID(gc.Request.Context(), id); err != nil {
		gc.JSON(http.StatusInternalServerError, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data: dto.ErrorData{
				Error: err.Error(),
			},
		})

		return
	}

	gc.JSON(http.StatusOK, nil)
}
