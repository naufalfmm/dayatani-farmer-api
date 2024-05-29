package persistents

import "github.com/naufalfmm/dayatani-farmer-api/persistents/repositories"

type Persistents struct {
	Repositories repositories.Repositories
}

func Init() (Persistents, error) {
	return Persistents{}, nil
}
