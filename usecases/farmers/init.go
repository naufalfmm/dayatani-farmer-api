package farmers

import "github.com/naufalfmm/dayatani-farmer-api/persistents"

type (
	Usecases interface{}

	usecases struct {
		persists persistents.Persistents
	}
)

func Init(persist persistents.Persistents) (Usecases, error) {
	return &usecases{
		persists: persist,
	}, nil
}
