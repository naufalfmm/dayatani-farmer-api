package farmers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

func (c Controllers) GetPaginated(gc *gin.Context) {
	var req dto.FarmerListPaginationRequest
	if err := req.FromGinContext(gc); err != nil {
		gc.JSON(http.StatusBadRequest, dto.Default{
			Ok:      false,
			Message: err.Error(),
			Data: dto.ErrorData{
				Error: err.Error(),
			},
		})

		return
	}

	pagingFarmer, err := c.Usecases.Farmers.GetPaginated(gc.Request.Context(), req)
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

	gc.JSON(http.StatusOK, dto.Default{
		Ok:      true,
		Message: "Success",
		Data:    dto.NewFarmerPaginationResponse(pagingFarmer),
	})
}
