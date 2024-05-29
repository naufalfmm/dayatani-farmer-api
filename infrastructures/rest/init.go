package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/infrastructures/rest/routes"
	"github.com/naufalfmm/dayatani-farmer-api/middlewares"
	"github.com/naufalfmm/dayatani-farmer-api/usecases"
)

type Rest struct {
	Routes routes.Routes
}

func Init(usc usecases.Usecases, midl middlewares.Middlewares) (Rest, error) {
	rout, err := routes.Init(usc, midl)
	if err != nil {
		return Rest{}, err
	}

	return Rest{
		Routes: rout,
	}, nil
}

func (r *Rest) Register(ge *gin.Engine) {
	r.Routes.Register(ge)
}
