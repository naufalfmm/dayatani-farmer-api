package farmers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naufalfmm/dayatani-farmer-api/persistents"
	"github.com/naufalfmm/dayatani-farmer-api/persistents/repositories"
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

		expUsec := usecases{
			persists: persist,
		}

		usec, err := Init(persist)

		assert.Nil(t, err)
		assert.Equal(t, &expUsec, usec)
	})
}

var (
	errAny = errors.New("any error")
)

type mock struct {
	ctrl *gomock.Controller
	ctx  context.Context

	farmer *mockFarmer.MockRepositories

	persistent persistents.Persistents

	usecases Usecases
}

func setupMock(t *testing.T) mock {
	mock := mock{}
	mock.ctrl = gomock.NewController(t)

	mock.farmer = mockFarmer.NewMockRepositories(mock.ctrl)

	mock.persistent = persistents.Persistents{
		Repositories: repositories.Repositories{
			Farmers: mock.farmer,
		},
	}

	mock.usecases = &usecases{
		persists: mock.persistent,
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	mock.ctx = req.Context()

	return mock
}

func (m *mock) Finish() {
	m.ctrl.Finish()
}
