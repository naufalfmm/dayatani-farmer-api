package farmers

import "github.com/naufalfmm/dayatani-farmer-api/usecases"

type Controllers struct {
	Usecases usecases.Usecases
}

func Init(usec usecases.Usecases) (Controllers, error) {
	return Controllers{
		Usecases: usec,
	}, nil
}
