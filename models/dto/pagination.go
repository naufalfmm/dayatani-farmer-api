package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/consts"
	"github.com/naufalfmm/dayatani-farmer-api/utils/helper"
)

type PaginationRequest struct {
	Page   int
	Limit  int
	Offset int
}

func (req *PaginationRequest) FromGinContext(gc *gin.Context) {
	req.Page = helper.ConvertStringDefault(gc.Param("page"), consts.DefPage)
	req.Limit = helper.ConvertStringDefault(gc.Param("limit"), consts.DefLimit)
	req.Offset = helper.ConvertStringDefault(gc.Param("offset"), consts.DefOffset)

	if req.Limit > consts.MaxLimit {
		req.Limit = consts.DefLimit
	}
}
