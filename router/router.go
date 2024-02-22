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
	public.GET("ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "pong!") })

	// 用户端操作
	public.POST("user/login", controller.UserLoginHandler)
	authed := public.Group("/", middleware.JWTUserAuth()) // 登录保护
	{
		authed.PUT("user/info", controller.UpdateUserInfoHandler)
		authed.GET("user/info", controller.GetUserInfoHandler)

		authed.GET("story/:id", controller.GetStoryHandler)
		authed.POST("story/save", controller.CreateStoryHandler)
		authed.POST("story/list", controller.ListStoryHandler)
		authed.DELETE("story/:id", controller.DeleteStoryHandler)
		authed.PUT("story", controller.UpdateStoryHandler)
	}

	// 管理端操作
	public.POST("admin/login", controller.AdminLoginHandler)
	admin := public.Group("/admin", middleware.JWTAdminAuth()) // 登录保护
	{
		admin.POST("register", controller.AdminRegisterHandler)
		admin.PUT("info", controller.UpdateAdminInfoHandler)
		admin.GET("info", controller.GetAdminInfoHandler)

		admin.GET("task/:id", controller.GetTaskHandler)
		admin.POST("task/save", controller.CreateTaskHandler)
		admin.POST("task/list", controller.ListTaskHandler)
		admin.DELETE("task/:id", controller.DeleteTaskHandler)
		admin.PUT("task", controller.UpdateTaskHandler)
	}

	return r
}
