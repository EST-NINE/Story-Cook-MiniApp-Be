package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/global"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
)

type OrderSrv struct {
}

func (s *OrderSrv) CreateOrder(ctx *gin.Context, req *dto.OrderDto) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	orderDao := dao.NewOrderDao(ctx)

	// 在开启新任务之后删除其他进行中的任务
	orders, err := orderDao.FindOngoingOrders(userInfo.Id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}
	for _, order := range orders {
		// 删除关联的故事
		err = dao.NewStoryDao(ctx).DeleteStory(order.StoryId)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}

		err = orderDao.DeleteOrder(order.ID)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
	}

	task, err := dao.NewTaskDao(ctx).FindTaskById(req.TaskId)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	story := dao.Story{
		UserId:  userInfo.Id,
		Title:   task.Title,
		Content: task.Content,
	}

	err = dao.NewStoryDao(ctx).CreateStory(&story)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	order := dao.Orders{
		UserId:  userInfo.Id,
		TaskId:  req.TaskId,
		StoryId: story.ID,
		Status:  global.OrderStatusOngoing, // 进行中
	}

	err = orderDao.CreateOrder(&order)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(vo.BuildOrderResp(&order)), nil
}

func (s *OrderSrv) FindOrderById(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	order, err := dao.NewOrderDao(ctx).FindOrderById(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(vo.BuildOrderResp(order)), nil
}

func (s *OrderSrv) DeleteOrder(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	orderDao := dao.NewOrderDao(ctx)

	order, err := orderDao.FindOrderById(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	// 删除关联的故事
	err = dao.NewStoryDao(ctx).DeleteStory(order.StoryId)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	err = orderDao.DeleteOrder(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

func (s *OrderSrv) UpdateOrder(ctx *gin.Context, req *dto.OrderDto) (resp *vo.Response, err error) {
	orderDao := dao.NewOrderDao(ctx)

	order := &dao.Orders{
		Comment: req.Comment,
		Score:   req.Score,
		Money:   req.Money,
		Status:  req.Status,
	}

	err = orderDao.UpdateOrder(req.ID, order)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

func (s *OrderSrv) ListOrder(ctx *gin.Context, req *dto.ListDto) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	orders, total, err := dao.NewOrderDao(ctx).ListOrder(req.Page, req.Limit, userInfo.Id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	listOrderResp := make([]*vo.OrderResp, 0)
	for _, order := range orders {
		listOrderResp = append(listOrderResp, vo.BuildOrderResp(order))
	}

	return vo.List(listOrderResp, total), nil
}

func (s *OrderSrv) SettleOrder(ctx *gin.Context, req *dto.OrderDto) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	// 生成货币的算法
	money := global.BasicOrderReward + req.Score/20

	// 更新订单信息，并标记为已完成
	order := &dao.Orders{
		Comment: req.Comment,
		Score:   req.Score,
		Money:   money,
		Status:  global.OrderStatusFinished, // 已完成
	}
	err = dao.NewOrderDao(ctx).UpdateOrder(req.ID, order)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	// 更新用户的货币
	err = dao.NewUserDao(ctx).AddReward(userInfo.Id, req.Money)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(money), nil
}
