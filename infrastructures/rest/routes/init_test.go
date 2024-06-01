package routes

import (
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/naufalfmm/dayatani-farmer-api/docs"
	"github.com/naufalfmm/dayatani-farmer-api/infrastructures/rest/controllers"
	"github.com/naufalfmm/dayatani-farmer-api/infrastructures/rest/controllers/farmers"
	mockMiddleware "github.com/naufalfmm/dayatani-farmer-api/mocks/middlewares"
	mockFarmer "github.com/naufalfmm/dayatani-farmer-api/mocks/usecases/farmers"
	"github.com/naufalfmm/dayatani-farmer-api/usecases"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	t.Run("If no error, it will return the routes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		farmer := mockFarmer.NewMockUsecases(ctrl)

		usecs := usecases.Usecases{
			Farmers: farmer,
		}

		farmerCont := farmers.Controllers{
			Usecases: usecs,
		}

		cont := controllers.Controllers{
			Farmers: farmerCont,
		}

		middl := mockMiddleware.NewMockMiddlewares(ctrl)

		expRout := Routes{
			Controllers: cont,
			Middlewares: middl,
		}

		rout, err := Init(usecs, middl)

		assert.Nil(t, err)
		assert.Equal(t, expRout, rout)
	})
}
