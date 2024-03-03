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
	user := public.Group("/user", middleware.JWTUserAuth()) // 登录保护
	{
		// 信息
		user.PUT("info", controller.UpdateUserInfoHandler)
		user.GET("info", controller.GetUserInfoHandler)

		// 故事
		user.GET("story/:id", controller.GetStoryHandler)
		user.POST("story/save", controller.CreateStoryHandler)
		user.POST("story/extend", controller.ExtendStoryHandler)
		user.POST("story/list", controller.ListStoryHandler)
		user.DELETE("story/:id", controller.DeleteStoryHandler)
		user.PUT("story", controller.UpdateStoryHandler)

		// 用户每日任务
		user.GET("task/:id", controller.GetUserTaskHandler)
		user.GET("tasks/:limit", controller.ListUserTaskHandler)
		user.POST("task/save", controller.CreateUserTaskHandler)
		user.DELETE("task/:id", controller.DeleteUserTaskHandler)
		user.PUT("task", controller.UpdateUserTaskHandler)
	}

	// 管理端操作
	public.POST("admin/login", controller.AdminLoginHandler)
	admin := public.Group("/admin", middleware.JWTAdminAuth()) // 登录保护
	{
		// 信息
		admin.POST("register", controller.AdminRegisterHandler)
		admin.PUT("info", controller.UpdateAdminInfoHandler)
		admin.GET("info", controller.GetAdminInfoHandler)

		// 每日任务
		admin.GET("task/:id", controller.GetTaskHandler)
		admin.POST("task/save", controller.CreateTaskHandler)
		admin.POST("task/list", controller.ListTaskHandler)
		admin.DELETE("task/:id", controller.DeleteTaskHandler)
		admin.PUT("task", controller.UpdateTaskHandler)
	}

	return r
}
