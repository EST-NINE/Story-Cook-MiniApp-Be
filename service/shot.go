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

	// 查询对应品质的菜品
	qualities := []string{"R", "SR", "SSR"}
	dishesMap, err := dishDao.ListDishesByQualities(qualities)
	if err != nil {
		return nil, err
	}

	var selectedDishes []*dao.Dish
	for i := 0; i < num; i++ {
		// 根据对应品质概率的权重抽取菜品
		randomNum := rand.Float64()
		var selectedDish *dao.Dish

		switch {
		case randomNum < global.ProbabilityR:
			selectedDish = dishesMap["R"][rand.Intn(len(dishesMap["R"]))]
		case randomNum < global.ProbabilityR+global.ProbabilitySR:
			selectedDish = dishesMap["SR"][rand.Intn(len(dishesMap["SR"]))]
		case randomNum < global.ProbabilityR+global.ProbabilitySR+global.ProbabilitySSR:
			selectedDish = dishesMap["SSR"][rand.Intn(len(dishesMap["SSR"]))]
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
	// 如果用户没有拥有过该菜品，则创建一条记录，数量置-1，标记为已解锁
	if errors.Is(err, gorm.ErrRecordNotFound) {
		userDish = &dao.UserDish{
			UserId:     userInfo.Id,
			DishId:     dish.ID,
			DishAmount: global.InitialUnlockDishAmount,
			IsUnlock:   global.InitialIsUnLock,
		}
		err := userDishDao.CreateUserDish(userDish)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
		return vo.SuccessWithData(vo.BuildShotResp(userDish, true)), nil
	}

	// 如果用户已经拥有过该菜品，则加碎片
	user.Piece += global.AddedPieceAmount
	err = userDao.UpdateUserById(userInfo.Id, user)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(vo.BuildShotResp(userDish, false)), nil
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
	userDishDao := dao.NewUserDishDao(ctx)
	dishes, err := generateRandomDish(ctx, 10)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	tenUserDishes := make([]*vo.ShotResp, 0)
	for i := 0; i < 10; i++ {
		dish := dishes[i]

		// 判断用户是否已经拥有该菜品
		userDish, err := userDishDao.FindUserDish(userInfo.Id, dish.ID)
		// 如果用户没有拥有过该菜品，则创建一条记录，数量置1，碎片置0
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userDish = &dao.UserDish{
				UserId:     userInfo.Id,
				DishId:     dish.ID,
				DishAmount: global.InitialUnlockDishAmount,
				IsUnlock:   global.InitialIsUnLock,
			}
			err := userDishDao.CreateUserDish(userDish)
			if err != nil {
				return vo.Error(err, myErrors.ErrorDatabase), err
			}

			// 将该菜品添加到返回结果中
			tenUserDishes = append(tenUserDishes, vo.BuildShotResp(userDish, true))
		} else {
			// 如果用户已经拥有过该菜品，则加碎片
			user.Piece += global.AddedPieceAmount
			err = userDao.UpdateUserById(userInfo.Id, user)
			if err != nil {
				return vo.Error(err, myErrors.ErrorDatabase), err
			}

			// 将该菜品添加到返回结果中
			tenUserDishes = append(tenUserDishes, vo.BuildShotResp(userDish, false))
		}
	}

	return vo.SuccessWithData(tenUserDishes), nil
}

func (s *ShotSrv) MergePiece(ctx *gin.Context, dishId uint) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserId(userInfo.Id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	// 判断碎片数量是否足够
	dish, err := dao.NewDishDao(ctx).FindDishById(dishId)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	// 根据菜品品质判断用户碎片数量是否足够
	switch dish.Quality {
	case "R":
		if user.Piece < global.PieceAmountR {
			err := errors.New("碎片数量不足")
			return vo.Error(err, myErrors.ErrorDatabase), err
		} else {
			user.Piece -= global.PieceAmountR
		}
	case "SR":
		if user.Piece < global.PieceAmountSR {
			err := errors.New("碎片数量不足")
			return vo.Error(err, myErrors.ErrorDatabase), err
		} else {
			user.Piece -= global.PieceAmountSR
		}
	case "SSR":
		if user.Piece < global.PieceAmountSSR {
			err := errors.New("碎片数量不足")
			return vo.Error(err, myErrors.ErrorDatabase), err
		} else {
			user.Piece -= global.PieceAmountSSR
		}
	}

	// 如果用户碎片足够，去判断用户有没有拥有这个菜品
	userDishDao := dao.NewUserDishDao(ctx)
	userDish, err := userDishDao.FindUserDish(userInfo.Id, dishId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果没有找到，创建新的用户菜品记录
		userDish = &dao.UserDish{
			UserId:     userInfo.Id,
			DishId:     dishId,
			DishAmount: 1,
			IsUnlock:   false,
		}

		err := userDishDao.CreateUserDish(userDish)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
		return vo.SuccessWithData(userDish), nil
	}

	// 如果有，则更新用户菜品记录
	userDish.DishAmount++
	err = userDishDao.UpdateUserDish(userDish)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	// 更新用户碎片数量
	err = userDao.UpdateUserById(userInfo.Id, user)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(userDish), nil
}
