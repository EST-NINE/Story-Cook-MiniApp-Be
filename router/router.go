package router

import (
	"net/http"

	"github.com/ncuhome/story-cook/config"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/controller"
	"github.com/ncuhome/story-cook/middleware"
)

func NewRouter() *gin.Engine {
	gin.SetMode(config.AppMode)
	r := gin.Default()

	r.Use(middleware.CORS())
	public := r.Group("api/v1")
	{
		public.GET("ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "pong!") })

		// 用户操作
		public.POST("user/login", controller.UserLoginHandler)

		authed := public.Group("/") // 登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user/info", controller.UpdateUserInfoHandler)
			authed.GET("user/info", controller.GetUserInfoHandler)

			// 故事操作
			authed.POST("story/save", controller.CreateStoryHandler)
			authed.POST("story/list", controller.ListStoryHandler)
			authed.DELETE("story/:id", controller.DeleteStoryHandler)
			authed.PUT("story", controller.UpdateStoryHandler)

			// 故事操作
			authed.POST("task/save", controller.CreateTaskHandler)
			authed.POST("task/list", controller.ListTaskHandler)
			authed.DELETE("task/:id", controller.DeleteTaskHandler)
			authed.PUT("task", controller.UpdateTaskHandler)
		}
	}

	return r
}
