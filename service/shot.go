package service

import (
	"errors"
	"math/rand"

	"github.com/ncuhome/story-cook/model/dto"

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
	for i := 0; i < num; {
		// 根据对应品质概率的权重抽取菜品
		randomNum := rand.Float64()
		var selectedDish *dao.Dish

		switch {
		case randomNum < global.ProbabilityMap["R"]:
			selectedDish = dishesMap["R"][rand.Intn(len(dishesMap["R"]))]
		case randomNum < global.ProbabilityMap["R"]+global.ProbabilityMap["SR"]:
			selectedDish = dishesMap["SR"][rand.Intn(len(dishesMap["SR"]))]
		default:
			selectedDish = dishesMap["SSR"][rand.Intn(len(dishesMap["SSR"]))]
		}

		if selectedDish != nil {
			selectedDishes = append(selectedDishes, selectedDish)
			i++
		}
	}

	return selectedDishes, nil
}

// SingleShot 单抽
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
	// 如果用户没有拥有过该菜品，则创建一条记录，数量置-1，标记为已解锁，并返回
	if errors.Is(err, gorm.ErrRecordNotFound) {
		userDish = &dao.UserDish{
			UserId:     userInfo.Id,
			DishId:     dish.ID,
			DishAmount: global.InitialUnlockDishAmount,
			IsUnlock:   true,
		}
		err := userDishDao.CreateUserDish(userDish)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
		return vo.SuccessWithData(vo.BuildShotResp(userDish, true, 0)), nil
	}

	// 找到了，但是如果用户未解锁这道菜品，则数量置-1，标记为已解锁，并返回
	if !userDish.IsUnlock {
		userDish.DishAmount = global.InitialUnlockDishAmount
		userDish.IsUnlock = true
		err := userDishDao.UpdateUserDish(userDish)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
		return vo.SuccessWithData(vo.BuildShotResp(userDish, true, 0)), nil
	}

	// 如果用户已经拥有过该菜品，则加对应品质的碎片以兑换其他的菜品
	user.Piece += global.PieceAmountMap[dish.Quality]
	err = userDao.UpdateUserById(userInfo.Id, user)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(vo.BuildShotResp(userDish, false, global.PieceAmountMap[dish.Quality])), nil
}

// TenShots 十连
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
		// 如果用户没有拥有过该菜品，则创建一条记录，数量置-1，标记为已解锁
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userDish = &dao.UserDish{
				UserId:     userInfo.Id,
				DishId:     dish.ID,
				DishAmount: global.InitialUnlockDishAmount,
				IsUnlock:   true,
			}
			err := userDishDao.CreateUserDish(userDish)
			if err != nil {
				return vo.Error(err, myErrors.ErrorDatabase), err
			}

			// 将该菜品添加到返回结果中
			tenUserDishes = append(tenUserDishes, vo.BuildShotResp(userDish, true, 0))
			continue
		}

		// 找到了，但是如果用户未解锁这道菜品，则数量置-1，标记为已解锁
		if !userDish.IsUnlock {
			userDish.DishAmount = global.InitialUnlockDishAmount
			userDish.IsUnlock = true
			err := userDishDao.UpdateUserDish(userDish)
			if err != nil {
				return vo.Error(err, myErrors.ErrorDatabase), err
			}

			// 将该菜品添加到返回结果中
			tenUserDishes = append(tenUserDishes, vo.BuildShotResp(userDish, true, 0))
			continue
		}

		// 如果用户已经拥有过该菜品，则加碎片
		user.Piece += global.PieceAmountMap[dish.Quality]

		// 将该菜品添加到返回结果中
		tenUserDishes = append(tenUserDishes, vo.BuildShotResp(userDish, false, global.PieceAmountMap[dish.Quality]))
	}

	// 更新用户信息
	err = userDao.UpdateUserById(userInfo.Id, user)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(tenUserDishes), nil
}

// MergePiece 合成未解锁的菜品
func (s *ShotSrv) MergePiece(ctx *gin.Context, req *dto.PieceDto) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserId(userInfo.Id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	// 查找菜品
	dish, err := dao.NewDishDao(ctx).FindDishById(req.DishId)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	// 根据菜品品质判断用户碎片数量是否足够
	requiredPieceAmount, ok := global.PieceAmountMap[dish.Quality]
	if !ok || user.Piece < requiredPieceAmount*req.ExchangeCount {
		err := errors.New("碎片数量不足")
		return vo.Error(err, myErrors.ErrorDatabase), err
	} else {
		user.Piece -= requiredPieceAmount * req.ExchangeCount
	}

	// 如果用户碎片足够，去判断用户有没有拥有这个菜品
	userDishDao := dao.NewUserDishDao(ctx)
	userDish, err := userDishDao.FindUserDish(userInfo.Id, req.DishId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果没有找到，创建新的用户菜品记录
		userDish = &dao.UserDish{
			UserId:     userInfo.Id,
			DishId:     req.DishId,
			DishAmount: req.ExchangeCount,
			IsUnlock:   false,
		}

		err := userDishDao.CreateUserDish(userDish)
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

	// 如果已拥有并且已解锁，返回
	if userDish.IsUnlock {
		err := errors.New("菜品已解锁")
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	// 如果已拥有并且未解锁，则更新用户菜品记录
	userDish.DishAmount += req.ExchangeCount
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
