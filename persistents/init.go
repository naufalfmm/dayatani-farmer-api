package persistents

import (
	"github.com/naufalfmm/dayatani-farmer-api/persistents/repositories"
	"github.com/naufalfmm/dayatani-farmer-api/resources/db"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger"
)

type Persistents struct {
	Repositories repositories.Repositories
}

func Init(d *db.DB, l logger.Logger) (Persistents, error) {
	repo, err := repositories.Init(d, l)
	if err != nil {
		return Persistents{}, err
	}

	return Persistents{
		Repositories: repo,
	}, nil
}
