package dto

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/consts"
	"github.com/naufalfmm/dayatani-farmer-api/utils/helper"
	"github.com/naufalfmm/dayatani-farmer-api/utils/orm"
)

type PaginationRequest struct {
	Limit  int
	Offset int
	Sorts  []string
}

func (req *PaginationRequest) FromGinContext(gc *gin.Context) {
	req.Limit = helper.ConvertStringDefault(gc.Query("limit"), consts.DefLimit)
	req.Offset = helper.ConvertStringDefault(gc.Query("offset"), consts.DefOffset)

	if gc.Query("sorts") != "" {
		req.Sorts = strings.Split(gc.Query("sorts"), ",")
	}

	if req.Limit > consts.MaxLimit {
		req.Limit = consts.DefLimit
	}
}

func (req *PaginationRequest) Paginate(selOrm orm.Select, mapSort map[string]func(ordKeyword string) string) orm.Select {
	for _, sort := range req.Sorts {
		ordKeyword := consts.OrdAsc
		if sort[0] == '-' {
			sort = sort[1:]
			ordKeyword = consts.OrdDesc
		}

		if _, ok := mapSort[sort]; !ok {
			continue
		}

		selOrm = selOrm.Order(mapSort[sort](ordKeyword))
	}

	selOrm = selOrm.Limit(req.Limit)

	return selOrm.Offset(req.Offset)
}

type PaginationResponse struct {
	Count  int      `json:"count"`
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
	Sorts  []string `json:"sorts"`
}
