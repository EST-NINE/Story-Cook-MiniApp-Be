package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
)

type DishSrv struct {
}

func (s *DishSrv) CreateDish(ctx *gin.Context, req *dto.DishDto) (resp *vo.Response, err error) {
	dish := &dao.Dish{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		Quality:     req.Quality,
	}

	err = dao.NewDishDao(ctx).CreateDish(dish)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(vo.BuildDishResp(dish)), nil
}

func (s *DishSrv) FindDishById(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	dish, err := dao.NewDishDao(ctx).FindDishById(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(vo.BuildDishResp(dish)), nil
}

func (s *DishSrv) DeleteDish(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	dishDao := dao.NewDishDao(ctx)

	// 删除关联的用户的食材表中的数据
	err = dao.NewUserDishDao(ctx).DeleteUserDishByDishId(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	// 删除对应的菜品存储在 oss 的图片
	dish, err := dishDao.FindDishById(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	err = util.DeleteFile(dish.Image)
	if err != nil {
		return vo.Error(err), err
	}

	// 删除对应的食材
	err = dishDao.DeleteDish(id)
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
		Quality:     req.Quality,
	}

	err = dishDao.UpdateDish(req.ID, dish)
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

	listDishResp := make([]*vo.DishResp, 0)
	for _, dish := range dishes {
		listDishResp = append(listDishResp, vo.BuildDishResp(dish))
	}
	return vo.List(listDishResp, total), nil
}

func (s *DishSrv) ListUserDish(ctx *gin.Context) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	userDishList, err := dao.NewUserDishDao(ctx).ListUserDish(userInfo.Id)
	return vo.List(userDishList, int64(len(userDishList))), nil
}
