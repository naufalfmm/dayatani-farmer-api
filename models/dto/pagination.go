package dto

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/consts"
	"github.com/naufalfmm/dayatani-farmer-api/utils/helper"
)

type PaginationRequest struct {
	Page   int
	Limit  int
	Offset int
	Sorts  []string
}

func (req *PaginationRequest) FromGinContext(gc *gin.Context) {
	req.Page = helper.ConvertStringDefault(gc.Param("page"), consts.DefPage)
	req.Limit = helper.ConvertStringDefault(gc.Param("limit"), consts.DefLimit)
	req.Offset = helper.ConvertStringDefault(gc.Param("offset"), consts.DefOffset)
	req.Sorts = strings.Split(gc.Param("sorts"), ",")

	if req.Limit > consts.MaxLimit {
		req.Limit = consts.DefLimit
	}
}
