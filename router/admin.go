package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/controller"
	"github.com/ncuhome/story-cook/middleware"
)

type AdminGroup struct{}

func NewAdminGroup() *AdminGroup {
	return &AdminGroup{}
}

func (a *AdminGroup) RegisterRoutes(group *gin.RouterGroup) {
	group.POST("login", controller.AdminLoginHandler)
	admin := group.Use(middleware.JWTAdminAuth()) // 登录保护
	{
		// 上传图片
		admin.POST("upload", controller.UploadImageHandler)

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

		// 食材
		admin.POST("dish/list", controller.ListDishHandler)
		admin.POST("dish/save", controller.CreateDishHandler)
		admin.DELETE("dish/:id", controller.DeleteDishHandler)
		admin.PUT("dish", controller.UpdateDishHandler)
		admin.GET("dish/:id", controller.GetDishHandler)

		// 用户
		admin.POST("user/list", controller.ListUserHandler)
		admin.DELETE("user/:id", controller.DeleteUserHandler)
		admin.PUT("user", controller.UpdateUserHandler)

		// Prompt
		admin.GET("prompt", controller.GetPromptHandler)
		admin.GET("prompt/list", controller.FindPromptListHandler)
		admin.PUT("prompt", controller.UpdatePromptHandler)
	}
}
