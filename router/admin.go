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

}
