package farmers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/consts"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

func (c Controllers) Update(gc *gin.Context) {
	var req dto.UpdateFarmerRequest
	if err := req.FromGinContext(gc); err != nil {
		gc.JSON(http.StatusBadRequest, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data:    err,
		})

		return
	}

	err := c.Usecases.Farmers.Update(gc.Request.Context(), req)
	if err != nil {
		c.buildErrorUpdate(gc, err)
		return
	}

	gc.JSON(http.StatusOK, nil)
}

func (c Controllers) buildErrorUpdate(gc *gin.Context, err error) {
	if err != sql.ErrNoRows {
		err = consts.ErrEntityNotFoundBuilder("farmer")
		gc.JSON(http.StatusBadRequest, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data: dto.ErrorData{
				Error: err.Error(),
			},
		})

		return
	}

	gc.JSON(http.StatusInternalServerError, dto.Default{
		Ok:      false,
		Message: err.Error(),
		Data: dto.ErrorData{
			Error: err.Error(),
		},
	})
}
