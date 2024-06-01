package farmers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

// Create Farmer godoc
//
//	@Summary		Create farmer
//	@Description	Create farmer
//	@Security		BasicAuth
//	@Tags			Farmers
//	@Accept			json
//	@Produce		json
//
//	@Param			farmer	body	dto.CreateFarmerRequest	true	"Farmer create request body"
//
//	@Success		201
//	@Failure		400	{object}	dto.Default{data=dto.ErrorData}
//	@Failure		500	{object}	dto.Default{data=dto.ErrorData}
//	@Router			/farmers [post]
func (c Controllers) Create(gc *gin.Context) {
	var req dto.CreateFarmerRequest
	if err := req.FromGinContext(gc); err != nil {
		gc.JSON(http.StatusBadRequest, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data:    err,
		})

		return
	}

	_, err := c.Usecases.Farmers.Create(gc.Request.Context(), req)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data: dto.ErrorData{
				Error: err.Error(),
			},
		})

		return
	}

	gc.Status(http.StatusCreated)
	gc.Writer.WriteHeaderNow()
}
