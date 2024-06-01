package usecases

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naufalfmm/dayatani-farmer-api/persistents"
	"github.com/naufalfmm/dayatani-farmer-api/persistents/repositories"
	"github.com/naufalfmm/dayatani-farmer-api/usecases/farmers"
	"github.com/stretchr/testify/assert"

	mockFarmer "github.com/naufalfmm/dayatani-farmer-api/mocks/persistents/repositories/farmers"
)

func Test_usecases_Init(t *testing.T) {
	t.Run("If no error, it will return the usecases", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		persist := persistents.Persistents{
			Repositories: repositories.Repositories{
				Farmers: mockFarmer.NewMockRepositories(ctrl),
			},
		}

		f, err := farmers.Init(persist)
		if err != nil {
			t.Error(err)
		}

		expUsec := Usecases{
			Farmers: f,
		}

		usec, err := Init(persist)

		assert.Nil(t, err)
		assert.Equal(t, expUsec, usec)
	})
}
