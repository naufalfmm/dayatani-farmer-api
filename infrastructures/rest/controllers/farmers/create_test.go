package farmers

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/naufalfmm/dayatani-farmer-api/models/dao"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
	"github.com/naufalfmm/dayatani-farmer-api/resources/validator"
	"github.com/stretchr/testify/assert"
)

func Test_Controllers_Create(t *testing.T) {
	t.Run("If the creating process is success, it will return 201", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		vld, err := validator.NewValidator()
		if err != nil {
			t.Error(err)
		}

		binding.Validator = vld

		bodyJsonReq := dto.CreateFarmerRequest{
			Name:          "Warga Kamboja",
			BirthDateBody: "1999-07-01",
		}

		mock.SetRequestBody(bodyJsonReq)

		bodyReq := dto.CreateFarmerRequest{
			Name:          bodyJsonReq.Name,
			BirthDateBody: bodyJsonReq.BirthDateBody,
			BirthDate:     time.Date(1999, 7, 1, 0, 0, 0, 0, time.UTC),
		}
		mock.farmer.EXPECT().Create(mock.ctx, bodyReq).Return(dao.Farmer{}, nil)

		mock.controllers.Create(mock.gc)

		assert.Equal(t, http.StatusCreated, mock.resRecorder.Code)
		assert.Equal(t, 0, mock.resRecorder.Body.Len())
	})

	t.Run("If the usecase returns error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		vld, err := validator.NewValidator()
		if err != nil {
			t.Error(err)
		}

		binding.Validator = vld

		bodyJsonReq := dto.CreateFarmerRequest{
			Name:          "Warga Kamboja",
			BirthDateBody: "1999-07-01",
		}

		mock.SetRequestBody(bodyJsonReq)

		bodyReq := dto.CreateFarmerRequest{
			Name:          bodyJsonReq.Name,
			BirthDateBody: bodyJsonReq.BirthDateBody,
			BirthDate:     time.Date(1999, 7, 1, 0, 0, 0, 0, time.UTC),
		}
		mock.farmer.EXPECT().Create(mock.ctx, bodyReq).Return(dao.Farmer{}, errAny)

		mock.controllers.Create(mock.gc)

		assert.Equal(t, http.StatusInternalServerError, mock.resRecorder.Code)
		assert.Equal(t, mock.MakeErrorResponse(errAny), mock.resRecorder.Body.String())
	})

	t.Run("If birth_date is missing, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		vld, err := validator.NewValidator()
		if err != nil {
			t.Error(err)
		}

		binding.Validator = vld

		bodyJsonReq := dto.CreateFarmerRequest{
			Name: "Warga Kamboja",
		}

		mock.SetRequestBody(bodyJsonReq)

		mock.controllers.Create(mock.gc)

		assert.Equal(t, http.StatusBadRequest, mock.resRecorder.Code)
	})

	t.Run("If the birth_date is wrong format, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		vld, err := validator.NewValidator()
		if err != nil {
			t.Error(err)
		}

		binding.Validator = vld

		bodyJsonReq := dto.CreateFarmerRequest{
			Name:          "Warga Kamboja",
			BirthDateBody: "01 Juli 1999",
		}

		mock.SetRequestBody(bodyJsonReq)

		mock.controllers.Create(mock.gc)

		assert.Equal(t, http.StatusBadRequest, mock.resRecorder.Code)
	})

	t.Run("If the name is missing, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		vld, err := validator.NewValidator()
		if err != nil {
			t.Error(err)
		}

		binding.Validator = vld

		bodyJsonReq := dto.CreateFarmerRequest{
			BirthDateBody: "1999-07-01",
		}

		mock.SetRequestBody(bodyJsonReq)

		mock.controllers.Create(mock.gc)

		assert.Equal(t, http.StatusBadRequest, mock.resRecorder.Code)
	})
}
