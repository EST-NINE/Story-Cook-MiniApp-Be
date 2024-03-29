package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
)

type OrderSrv struct {
}

func (s *OrderSrv) CreateOrder(ctx *gin.Context, req *dto.OrderDto) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	order := dao.Orders{
		UserId:  userInfo.Id,
		TaskId:  req.TaskId,
		StoryId: req.StoryId,
		Status:  1, // 进行中
	}

	err = dao.NewOrderDao(ctx).CreateOrder(&order)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(order), nil
}

func (s *OrderSrv) FindOrderById(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	order, err := dao.NewOrderDao(ctx).FindOrderById(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(order), nil
}

func (s *OrderSrv) DeleteOrder(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	err = dao.NewOrderDao(ctx).DeleteOrder(id)
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

	return vo.List(orders, total), nil
}
