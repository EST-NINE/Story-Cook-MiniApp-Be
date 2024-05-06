package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/controller"
	"github.com/ncuhome/story-cook/middleware"
)

type UserGroup struct{}

func NewUserGroup() *UserGroup {
	return &UserGroup{}
}

func (u *UserGroup) RegisterRoutes(group *gin.RouterGroup) {
	group.POST("login", controller.UserLoginHandler)
	user := group.Use(middleware.JWTUserAuth()) // 登录保护
	{
		// 信息
		user.PUT("info", controller.UpdateUserInfoHandler)
		user.GET("info", controller.GetUserInfoHandler)

		// 故事
		user.GET("story/:id", controller.GetStoryHandler)
		user.POST("story/save", controller.CreateStoryHandler)
		user.POST("story/extend", controller.ExtendStoryHandler)
		user.POST("story/end", controller.EndStoryHandler)
		user.POST("story/assess/:id", controller.AssessStoryHandler)
		user.POST("story/list", controller.ListStoryHandler)
		user.DELETE("story/:id", controller.DeleteStoryHandler)
		user.PUT("story", controller.UpdateStoryHandler)

		// 每日任务
		user.GET("task", controller.GetDailyTaskHandler)
		user.GET("task/ongoing", controller.GetOngoingHandler)
		user.POST("task/list", controller.ListUserTaskHandler)

		// 订单
		user.GET("order/:id", controller.GetOrderHandler)
		user.POST("orders", controller.ListOrderHandler)
		user.POST("order/save", controller.CreateOrderHandler)
		user.DELETE("order/:id", controller.DeleteOrderHandler)
		user.PUT("order", controller.UpdateOrderHandler)
		user.POST("order/settle", controller.SettleOrderHandler)

		// 食材
		user.GET("dish", controller.ListUserDishHandler)

		// 抽卡
		user.GET("shot/single", controller.ShotSingleHandler)
		user.GET("shot/ten", controller.TenShotsHandler)
		user.POST("shot/merge", controller.MergePieceHandler)
	}
}
