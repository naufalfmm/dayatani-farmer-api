package infrastructures

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/infrastructures/rest"
	"github.com/naufalfmm/dayatani-farmer-api/middlewares"
	"github.com/naufalfmm/dayatani-farmer-api/usecases"
)

type Infrastructures struct {
	Rest rest.Rest
}

func Init(usc usecases.Usecases, midl middlewares.Middlewares) (Infrastructures, error) {
	rs, err := rest.Init(usc, midl)
	if err != nil {
		return Infrastructures{}, err
	}

	return Infrastructures{
		Rest: rs,
	}, nil
}

func (i *Infrastructures) Register(ge *gin.Engine) {
	i.Rest.Register(ge)
}
