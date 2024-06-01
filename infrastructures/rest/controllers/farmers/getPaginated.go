package farmers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

// Get Paginated Farmers godoc
//
//	@Summary		Get paginated farmers
//	@Description	Get paginated farmers
//	@Security		BasicAuth
//	@Tags			Farmers
//	@Accept			json
//	@Produce		json
//
//	@Param			limit	query		int		false	"The number of returned farms"
//	@Param			offset	query		int		false	"The number of skip specified rows"
//	@Param			sorts	query		string	false	"The sort separated by comma with "-" if the sort is descending"
//
//	@Success		200		{object}	dto.Default{data=dto.FarmerPaginationResponse}
//	@Failure		400		{object}	dto.Default{data=dto.ErrorData}
//	@Failure		500		{object}	dto.Default{data=dto.ErrorData}
//	@Router			/farmers [get]
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
