package router

import (
	"net/http"

	"github.com/ncuhome/story-cook/config"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/middleware"
)

func NewRouter() *gin.Engine {
	gin.SetMode(config.AppMode)
	r := gin.Default()

	r.Use(middleware.CORS())
	public := r.Group("api/v1")
	public.GET("ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "pong!") })

	// 用户端
	NewUserGroup().RegisterRoutes(public.Group("user"))

	// 管理端
	NewAdminGroup().RegisterRoutes(public.Group("admin"))

	return r
}
