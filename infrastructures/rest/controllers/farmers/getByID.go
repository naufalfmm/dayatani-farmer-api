package farmers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/consts"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

// Get Farmer by ID godoc
//
//	@Summary		Get farmer detail by id
//	@Description	Get farmer detail by id
//	@Security		BasicAuth
//	@Tags			Farmers
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		uint64	true	"Farmer id"
//
//	@Success		200	{object}	dto.Default{data=dto.FarmerResponse}
//	@Failure		400	{object}	dto.Default{data=dto.ErrorData}
//	@Failure		500	{object}	dto.Default{data=dto.ErrorData}
//	@Router			/farmers/{id} [get]
func (c Controllers) GetByID(gc *gin.Context) {
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

	frm, err := c.Usecases.Farmers.GetByID(gc.Request.Context(), id)
	if err != nil {
		c.buildErrorGetByID(gc, err)
		return
	}

	gc.JSON(http.StatusOK, dto.Default{
		Ok:      true,
		Message: "Success",
		Data:    dto.NewFarmerResponse(frm),
	})
}

func (c Controllers) buildErrorGetByID(gc *gin.Context, err error) {
	if err == sql.ErrNoRows {
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
