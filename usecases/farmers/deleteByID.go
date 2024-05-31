package farmers

import (
	"context"
)

func (u usecases) DeleteByID(ctx context.Context, id uint64) error {
	return u.persists.Repositories.Farmers.DeleteByID(ctx, id)
}
