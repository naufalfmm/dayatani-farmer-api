package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/dayatani-farmer-api/consts"
	"github.com/naufalfmm/dayatani-farmer-api/models/dto"
)

func (m middlewares) VerifyAuth() gin.HandlerFunc {
	return func(gc *gin.Context) {
		tokenAuth := gc.Request.Header.Get("Authorization")
		if tokenAuth == "" {
			gc.AbortWithStatusJSON(http.StatusUnauthorized, dto.Default{
				Ok:      false,
				Message: consts.ErrInvalidAuth.Error(),
				Data: dto.ErrorData{
					Error: consts.ErrInvalidAuth.Error(),
				},
			})
			return
		}

		token := ""
		tokenSplit := strings.Split(tokenAuth, " ")
		if len(tokenSplit) > 1 {
			token = tokenSplit[1]
		}

		decodedToken, err := m.e.Decode(token)
		if err != nil {
			gc.AbortWithStatusJSON(http.StatusUnauthorized, dto.Default{
				Ok:      false,
				Message: consts.ErrInvalidAuth.Error(),
				Data: dto.ErrorData{
					Error: consts.ErrInvalidAuth.Error(),
				},
			})
			return
		}

		spldToken := strings.Split(decodedToken, ":")

		if err := m.h.Check(m.c.HashedAuthUsername, spldToken[0]); err != nil {
			gc.AbortWithStatusJSON(http.StatusUnauthorized, dto.Default{
				Ok:      false,
				Message: consts.ErrInvalidAuth.Error(),
				Data: dto.ErrorData{
					Error: consts.ErrInvalidAuth.Error(),
				},
			})
			return
		}

		if err := m.h.Check(m.c.HashedAuthPassword, spldToken[1]); err != nil {
			gc.AbortWithStatusJSON(http.StatusUnauthorized, dto.Default{
				Ok:      false,
				Message: consts.ErrInvalidAuth.Error(),
				Data: dto.ErrorData{
					Error: consts.ErrInvalidAuth.Error(),
				},
			})
			return
		}

		gc.Next()

	}
}
