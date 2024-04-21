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

// 随机生成指定数量的菜品
func generateRandomDish(ctx *gin.Context, num int) ([]*dao.Dish, error) {
	dishDao := dao.NewDishDao(ctx)
	var selectedDishes []*dao.Dish

	// 查询对应品质的菜品
	dishesWithR, totalWithR, err := dishDao.ListDishByQuality("R")
	if err != nil {
		return nil, err
	}

	dishesWithSR, totalWithSR, err := dishDao.ListDishByQuality("SR")
	if err != nil {
		return nil, err
	}

	dishesWithSSR, totalWithSSR, err := dishDao.ListDishByQuality("SSR")
	if err != nil {
		return nil, err
	}

	// 根据对应品质概率的权重抽取菜品
	for i := 0; i < num; i++ {
		randomNum := rand.Float64()
		var selectedDish *dao.Dish

		switch {
		case randomNum < global.ProbabilityR:
			selectedDish = dishesWithR[rand.Intn(int(totalWithR))]
		case randomNum < global.ProbabilityR+global.ProbabilitySR:
			selectedDish = dishesWithSR[rand.Intn(int(totalWithSR))]
		case randomNum < global.ProbabilityR+global.ProbabilitySR+global.ProbabilitySSR:
			selectedDish = dishesWithSSR[rand.Intn(int(totalWithSSR))]
		}

		selectedDishes = append(selectedDishes, selectedDish)
	}

	return selectedDishes, nil
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

	if user.Money < global.SingleShotCost {
		err = errors.New("代币不足")
		return vo.Error(err, myErrors.ErrorDatabase), err
	} else {
		user.Money -= global.SingleShotCost
		err = userDao.UpdateUserById(userInfo.Id, user)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
	}

	// 随机生成一个菜品
	dishes, err := generateRandomDish(ctx, 1)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}
	dish := dishes[0]

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

func (s *ShotSrv) TenShots(ctx *gin.Context) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	// 判断抽卡的代币是否足够
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserId(userInfo.Id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	if user.Money < global.TenShotCost {
		err = errors.New("代币不足")
		return vo.Error(err, myErrors.ErrorDatabase), err
	} else {
		user.Money -= global.TenShotCost
		err = userDao.UpdateUserById(userInfo.Id, user)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
	}

	// 随机生成十个菜品（可重复）
	dishes, total, err := dao.NewDishDao(ctx).ListDish()
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	tenUserDishes := make([]*dao.UserDish, 0)
	selectedDishIDs := make(map[uint]bool)
	for i := 0; i < 10; i++ {
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

			// 将该菜品添加到返回结果中
			if !selectedDishIDs[dish.ID] {
				tenUserDishes = append(tenUserDishes, userDish)
			}
			selectedDishIDs[dish.ID] = true
		} else {
			// 如果用户已经拥有过该菜品，则加碎片
			userDish.PieceAmount += global.AddedPieceAmount
			err = userDishDao.UpdateUserDish(userDish)
			if err != nil {
				return vo.Error(err, myErrors.ErrorDatabase), err
			}

			// 将该菜品添加到返回结果中
			if !selectedDishIDs[dish.ID] {
				tenUserDishes = append(tenUserDishes, userDish)
			}
			selectedDishIDs[dish.ID] = true
		}
	}

	return vo.SuccessWithData(tenUserDishes), nil
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
