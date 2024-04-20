package service

import (
	"errors"
	"math/rand"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/global"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
)

type ShotSrv struct {
}

func (s *ShotSrv) SingleShot(ctx *gin.Context) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	// 判断抽卡的代币是否足够
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserId(userInfo.Id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	if user.Money < global.SingleShotMoney {
		err = errors.New("代币不足")
		return vo.Error(err, myErrors.ErrorDatabase), err
	} else {
		user.Money -= global.SingleShotMoney
		err = userDao.UpdateUserById(userInfo.Id, user)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
	}

	// 随机生成一个菜品
	dishes, total, err := dao.NewDishDao(ctx).ListDish()
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}
	random := rand.Intn(int(total))
	dish := dishes[random]

	// 判断用户是否已经拥有该菜品
	userDishDao := dao.NewUserDishDao(ctx)
	userDish, err := userDishDao.FindUserDish(userInfo.Id, dish.ID)
	// 如果用户没有拥有过该菜品，则创建一条记录，数量置1，碎片置0
	if errors.Is(err, gorm.ErrRecordNotFound) {
		userDish = &dao.UserDish{
			UserId:      userInfo.Id,
			DishId:      dish.ID,
			DishAmount:  global.InitialDishAmount,
			PieceAmount: global.InitialPieceAmount,
		}
		err := userDishDao.CreateUserDish(userDish)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
		return vo.SuccessWithData(userDish), nil
	}

	// 如果用户已经拥有过该菜品，则加碎片
	userDish.PieceAmount += global.AddedPieceAmount
	err = userDishDao.UpdateUserDish(userDish)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(userDish), nil
}

func (s *ShotSrv) MergePiece(ctx *gin.Context, dishId uint) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	userDishDao := dao.NewUserDishDao(ctx)
	userDish, err := userDishDao.FindUserDish(userInfo.Id, dishId)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	if userDish.PieceAmount < global.MergePieceAmount {
		err := errors.New("碎片数量不足")
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	userDish.PieceAmount -= global.MergePieceAmount
	userDish.DishAmount++
	err = userDishDao.UpdateUserDish(userDish)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(userDish), nil
}
