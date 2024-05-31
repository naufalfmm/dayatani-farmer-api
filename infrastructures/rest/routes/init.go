package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/infrastructures/rest/controllers"
	"github.com/naufalfmm/dayatani-farmer-api/middlewares"
	"github.com/naufalfmm/dayatani-farmer-api/usecases"
)

type Routes struct {
	Controllers controllers.Controllers
	Middlewares middlewares.Middlewares
}

func Init(usc usecases.Usecases, midl middlewares.Middlewares) (Routes, error) {
	c, err := controllers.Init(usc)
	if err != nil {
		return Routes{}, err
	}

	return Routes{
		Controllers: c,
		Middlewares: midl,
	}, nil
}

func (r Routes) Register(ge *gin.Engine) {
	ge.RedirectFixedPath = true

	fr := ge.Group("/farmers")
	fr.GET("/:id", r.Controllers.Farmers.GetByID)
	fr.GET("", r.Controllers.Farmers.GetPaginated)
	fr.POST("", r.Controllers.Farmers.Create)
	fr.PUT("/:id", r.Controllers.Farmers.Update)
}
