package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/consts"
	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
)

type FarmerResponse struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
}

func NewFarmerResponse(f dao.Farmer) FarmerResponse {
	return FarmerResponse{
		ID:        f.ID,
		Name:      f.Name,
		BirthDate: f.BirthDate.Format(consts.LayoutISOTime),
	}
}

type FarmerListPaginationRequest struct {
	PaginationRequest
}

func (req *FarmerListPaginationRequest) FromGinContext(gc *gin.Context) error {
	req.PaginationRequest.FromGinContext(gc)

	if len(req.Sorts) == 0 {
		req.Sorts = []string{"name"}
	}

	return nil
}

type FarmerPaginationResponse struct {
	PaginationResponse
	Items []FarmerResponse `json:"items"`
}

func NewFarmerPaginationResponse(fp dao.FarmerPaging) FarmerPaginationResponse {
	frs := make([]FarmerResponse, len(fp.Items))
	for i, item := range fp.Items {
		frs[i] = NewFarmerResponse(item)
	}

	return FarmerPaginationResponse{
		PaginationResponse: PaginationResponse{
			Count:  fp.Count,
			Limit:  fp.Limit,
			Offset: fp.Offset,
			Sorts:  fp.Sorts,
		},
		Items: frs,
	}
}
