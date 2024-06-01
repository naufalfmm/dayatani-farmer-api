package farmers

import (
	"context"
	"fmt"

	"github.com/naufalfmm/dayatani-farmer-api/consts"
	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

func (r repositories) GetPaginated(ctx context.Context, req dto.FarmerListPaginationRequest) (dao.FarmerPaging, error) {
	mapSort := map[string]func(ordKeyword string) string{
		"name": func(ordKeyword string) string {
			return fmt.Sprintf("%s %s", ColumnFarmerName, ordKeyword)
		},
	}

	fs := dao.Farmers{}
	selOrm := req.Paginate(r.db.GetDB().NewSelect().Model(&fs), mapSort)

	count, err := selOrm.ScanAndCount(ctx)
	if err != nil {
		r.log.Error(ctx, LogMsgGetPaginatedFarmers).Err(err).Any(consts.KeyReq, req).Send()
		return dao.FarmerPaging{}, err
	}

	return dao.FarmerPaging{
		Offset: req.Offset,
		Limit:  req.Limit,
		Count:  count,
		Sorts:  req.Sorts,

		Items: fs,
	}, nil
}
