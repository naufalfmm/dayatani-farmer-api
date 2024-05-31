package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/resources/config"
	"github.com/naufalfmm/dayatani-farmer-api/utils/encoding"
	"github.com/naufalfmm/dayatani-farmer-api/utils/hashing"
)

type (
	Middlewares interface {
		VerifyAuth() gin.HandlerFunc
		PanicRecover() gin.HandlerFunc
		ImplementCors() gin.HandlerFunc
	}

	middlewares struct {
		h hashing.Hashing
		e encoding.Encoding
		c *config.EnvConfig
	}
)

func Init(h hashing.Hashing, e encoding.Encoding, c *config.EnvConfig) (Middlewares, error) {
	return &middlewares{
		h: h,
		e: e,
		c: c,
	}, nil
}
