package farmers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

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

	_, err := c.Usecases.Farmers.Create(gc.Request.Context(), req.ToFarmer())
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

	gc.JSON(http.StatusCreated, nil)
}
