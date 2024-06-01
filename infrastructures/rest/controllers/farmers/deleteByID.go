package farmers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

// Delete Farmer by ID godoc
//
//	@Summary		Delete farmer detail by id
//	@Description	Delete farmer detail by id
//	@Security		BasicAuth
//	@Tags			Farmers
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path	uint64	true	"Farmer id"
//
//	@Success		200
//	@Failure		400	{object}	dto.Default{data=dto.ErrorData}
//	@Failure		500 {object}	dto.Default{data=dto.ErrorData}
//	@Router			/farmers/{id} [delete]
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

	gc.Status(http.StatusOK)
	gc.Writer.WriteHeaderNow()
}
