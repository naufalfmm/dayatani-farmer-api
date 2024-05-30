package persistents

import (
	"github.com/naufalfmm/dayatani-farmer-api/persistents/repositories"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
	"github.com/naufalfmm/dayatani-farmer-api/utils/orm"
)

type Persistents struct {
	Repositories repositories.Repositories
}

func Init(o orm.Orm, l logger.Logger) (Persistents, error) {
	repo, err := repositories.Init(o, l)
	if err != nil {
		return Persistents{}, err
	}

	return Persistents{
		Repositories: repo,
	}, nil
}
