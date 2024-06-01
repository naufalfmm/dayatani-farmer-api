package rest

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naufalfmm/dayatani-farmer-api/infrastructures/rest/controllers"
	"github.com/naufalfmm/dayatani-farmer-api/infrastructures/rest/controllers/farmers"
	"github.com/naufalfmm/dayatani-farmer-api/infrastructures/rest/routes"
	mockMiddleware "github.com/naufalfmm/dayatani-farmer-api/mocks/middlewares"
	mockFarmer "github.com/naufalfmm/dayatani-farmer-api/mocks/usecases/farmers"
	"github.com/naufalfmm/dayatani-farmer-api/usecases"
	"github.com/stretchr/testify/assert"
)

func Test_Rest_Init(t *testing.T) {
	t.Run("If no error, it will return the rest", func(t *testing.T) {
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

		rout := routes.Routes{
			Controllers: cont,
			Middlewares: middl,
		}

		expRe := Rest{
			Routes: rout,
		}

		re, err := Init(usecs, middl)

		assert.Nil(t, err)
		assert.Equal(t, expRe, re)
	})
}
