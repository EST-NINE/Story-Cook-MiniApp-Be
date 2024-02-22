package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
)

func JWTUserAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := myErrors.SUCCESS
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = myErrors.ErrorAuthCheckTokenFail
			ctx.JSON(http.StatusBadRequest, vo.Error(errors.New("empty token"), code))
			ctx.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil || claims.Authority != 0 {
			code = myErrors.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = myErrors.ErrorAuthCheckTokenTimeout
		}

		if code != myErrors.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, vo.Error(errors.New("token incorrect"), code))
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}

func JWTAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := myErrors.SUCCESS
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = myErrors.ErrorAuthCheckTokenFail
			ctx.JSON(http.StatusBadRequest, vo.Error(errors.New("empty token"), code))
			ctx.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil || claims.Authority != 1 {
			code = myErrors.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = myErrors.ErrorAuthCheckTokenTimeout
		}

		if code != myErrors.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, vo.Error(errors.New("token incorrect"), code))
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
