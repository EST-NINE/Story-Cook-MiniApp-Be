package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
)

type DishSrv struct {
}

func (s *DishSrv) CreateDish(ctx *gin.Context, req *dto.DishDto) (resp *vo.Response, err error) {
	dish := dao.Dish{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
	}

	err = dao.NewDishDao(ctx).CreateDish(&dish)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(dish), nil
}

func (s *DishSrv) FindDishById(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	dish, err := dao.NewDishDao(ctx).FindDishById(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(dish), nil
}

func (s *DishSrv) DeleteDish(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	err = dao.NewDishDao(ctx).DeleteDish(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

func (s *DishSrv) UpdateDish(ctx *gin.Context, req *dto.DishDto) (resp *vo.Response, err error) {
	dishDao := dao.NewDishDao(ctx)

	dish := &dao.Dish{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
	}

	err = dishDao.UpdateDish(req.ID, dish)
	fmt.Println(dish)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

func (s *DishSrv) ListDish(ctx *gin.Context) (resp *vo.Response, err error) {
	dishes, total, err := dao.NewDishDao(ctx).ListDish()
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.List(dishes, total), nil
}
