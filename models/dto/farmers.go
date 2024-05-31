package dto

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/consts"
	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/utils/helper"
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

type UpsertFarmerRequest struct {
	Name          string `json:"name" validate:"required"`
	BirthDateBody string `json:"birth_date" validate:"required,datetime=2006-01-02"`

	BirthDate time.Time `json:"-"`
}

func (req *UpsertFarmerRequest) FromGinContext(gc *gin.Context) error {
	if err := gc.ShouldBindJSON(req); err != nil {
		return helper.HandleBindError(*req, err)
	}

	fmtd, err := time.Parse(consts.LayoutISOTime, req.BirthDateBody)
	if err != nil {
		return err
	}

	req.BirthDate = fmtd

	return nil
}

func (req UpsertFarmerRequest) ToFarmer() dao.Farmer {
	return dao.Farmer{
		Name:      req.Name,
		BirthDate: req.BirthDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
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
