package repositories

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/naufalfmm/dayatani-farmer-api/persistents/repositories/farmers"
	"github.com/naufalfmm/dayatani-farmer-api/resources/db"
	"github.com/naufalfmm/dayatani-farmer-api/utils/logger/mockLogger"
	"github.com/naufalfmm/dayatani-farmer-api/utils/orm/mockOrm"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_Init(t *testing.T) {
	t.Run("If no error, it will return the repositories", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orm := mockOrm.NewMockOrm(ctrl)
		log := mockLogger.NewMockLogger(ctrl)

		db := db.DB{
			O: orm,
		}

		f, err := farmers.Init(&db, log)
		if err != nil {
			t.Error(err)
		}

		expRepo := Repositories{
			Farmers: f,
		}

		repo, err := Init(&db, log)

		assert.Nil(t, err)
		assert.Equal(t, expRepo, repo)
	})
}
