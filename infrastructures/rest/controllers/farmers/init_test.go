package farmers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/naufalfmm/dayatani-farmer-api/consts"
	mockFarmer "github.com/naufalfmm/dayatani-farmer-api/mocks/usecases/farmers"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
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

		res, err := Init(usecs)

		assert.Nil(t, err)
		assert.Equal(t, Controllers{Usecases: usecs}, res)
	})
}

var (
	errAny = errors.New("any error")
)

type mock struct {
	ctrl        *gomock.Controller
	ctx         context.Context
	gc          *gin.Context
	resRecorder *httptest.ResponseRecorder

	farmer *mockFarmer.MockUsecases

	controllers Controllers
}

func setupMock(t *testing.T) mock {
	mock := mock{}
	mock.ctrl = gomock.NewController(t)

	mock.farmer = mockFarmer.NewMockUsecases(mock.ctrl)

	mock.controllers = Controllers{
		Usecases: usecases.Usecases{
			Farmers: mock.farmer,
		},
	}

	mock.resRecorder = httptest.NewRecorder()

	mock.gc, _ = gin.CreateTestContext(mock.resRecorder)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	mock.gc.Request = req

	mock.ctx = mock.gc.Request.Context()

	return mock
}

func (m mock) Finish() {
	m.ctrl.Finish()
}

func (m mock) SetRequestBody(body interface{}) {
	requestByte, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(requestByte)))

	m.gc.Request = req
	m.gc.Request.Header.Set(consts.ContentTypeHeader, consts.ApplicationJsonContentType)

	m.ctx = req.Context()
}

func (m mock) SetParam(key, val string) {
	m.gc.Params = append(m.gc.Params, gin.Param{
		Key:   key,
		Value: val,
	})
}

func (m mock) SetURL(u string) {
	if m.gc.Request == nil {
		m.gc.Request = httptest.NewRequest(http.MethodGet, u, nil)
	}

	m.gc.Request.URL, _ = url.Parse(u)
}

func (m mock) makeDataResponse(expectedContent interface{}) string {
	expectedByte, _ := json.Marshal(expectedContent)
	expectedResult := string(expectedByte)
	return expectedResult
}

func (m mock) MakeSuccessResponse(message string, expContent interface{}) string {
	expResp := dto.Default{
		Ok:      true,
		Message: "Success",
		Data:    expContent,
	}

	return m.makeDataResponse(expResp)
}

func (m mock) MakeErrorResponse(err error) string {
	expResp := dto.Default{
		Ok:      false,
		Message: err.Error(),
		Data: dto.ErrorData{
			Error: err.Error(),
		},
	}

	return m.makeDataResponse(expResp)
}

func (m mock) MakeValidateErrorResponse(err error) string {
	expResp := dto.Default{
		Ok:      false,
		Message: err.Error(),
		Data:    err,
	}

	return m.makeDataResponse(expResp)
}
