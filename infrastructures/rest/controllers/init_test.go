package controllers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naufalfmm/dayatani-farmer-api/infrastructures/rest/controllers/farmers"
	mockFarmer "github.com/naufalfmm/dayatani-farmer-api/mocks/usecases/farmers"
	"github.com/naufalfmm/dayatani-farmer-api/usecases"
	"github.com/stretchr/testify/assert"
)

func Test_Controllers_Init(t *testing.T) {
	t.Run("If no error, it will return the controllers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		farmer := mockFarmer.NewMockUsecases(ctrl)

		usecs := usecases.Usecases{
			Farmers: farmer,
		}

		farmerCont := farmers.Controllers{
			Usecases: usecs,
		}

		expContrs := Controllers{
			Farmers: farmerCont,
		}

		cont, err := Init(usecs)

		assert.Nil(t, err)
		assert.Equal(t, expContrs, cont)
	})
}
