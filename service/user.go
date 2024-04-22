package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ncuhome/story-cook/pkg/myErrors"

	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/pkg/util"
	"gorm.io/gorm"
)

type UserSrv struct {
}

// Login 用户登陆函数
func (s *UserSrv) Login(ctx *gin.Context, req *dto.UserDto) (resp *vo.Response, err error) {

	// 通过 code 获取用户openid
	openid, err := util.GetWxOpenid(req.Code)
	if openid == "" || err != nil {
		err = errors.New("wxCode incorrect")
		return vo.Error(err, myErrors.ErrorInvalidParams), err
	}

	userDao := dao.NewUserDao(ctx)
	loginDao := dao.NewDailyLoginDao(ctx)

	// 查询数据库，判断用户是否存在，没找到则创建一个用户
	user, err := userDao.FindUserByOpenid(openid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &dao.User{
			UserName: fmt.Sprint("用户" + uuid.New().String()[:6]),
			Openid:   openid,
		}
		if err = userDao.CreateUser(user); err != nil {
			return vo.Error(err, myErrors.ErrorCreateUser), err
		}

		if err = loginDao.CreateDailyLogin(&dao.DailyLogin{
			UserId: user.ID,
			Date:   time.Now(),
		}); err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
	}

	// 实现每日签到的奖励
	if _, err = loginDao.FindDailyLoginById(user.ID); errors.Is(err, gorm.ErrRecordNotFound) {
		if err = userDao.DailyLoginReward(user); err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}

		user, err = userDao.FindUserByOpenid(openid)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
	}

	token, err := util.GenerateToken(user.ID, 0)
	if err != nil {
		return vo.Error(err, myErrors.ErrorAuthToken), err
	}
	return vo.SuccessWithDataAndToken(user, token), nil
}

// UserInfo 得到用户的信息
func (s *UserSrv) UserInfo(ctx *gin.Context) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserId(userInfo.Id)
	if err != nil {
		return vo.Error(err), err
	}

	respData := vo.BuildUserResp(user)
	return vo.SuccessWithData(respData), nil
}

// UpdateInfo 用户更改信息
func (s *UserSrv) UpdateInfo(ctx *gin.Context, req *dto.UserDto) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserId(userInfo.Id)
	if req.UserName != "" {
		user.UserName = req.UserName
	}

	if user.Money+req.Money < 0 {
		err = errors.New("money not enough")
		return vo.Error(err, myErrors.ErrorNotEnoughMoney), err
	} else {
		user.Money += req.Money
	}

	if user.Piece+req.Piece < 0 {
		err = errors.New("piece not enough")
		return vo.Error(err, myErrors.ErrorNotEnoughMoney), err
	} else {
		user.Piece += req.Piece
	}

	err = userDao.UpdateUserById(userInfo.Id, user)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	userResp := vo.BuildUserResp(user)
	return vo.SuccessWithData(userResp), nil
}

func (s *UserSrv) ListUser(ctx *gin.Context, req *dto.ListUserDto) (resp *vo.Response, err error) {
	userDao := dao.NewUserDao(ctx)
	var users []*dao.User
	var total int64

	switch req.Order {
	case 0:
		users, total, err = userDao.ListUserByID(req.Page, req.Limit)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
	case 1:
		users, total, err = userDao.ListUserByMoney(req.Page, req.Limit)
		if err != nil {
			return vo.Error(err, myErrors.ErrorDatabase), err
		}
	}

	listUserResp := make([]*vo.UserResp, 0)
	for _, task := range users {
		listUserResp = append(listUserResp, vo.BuildUserResp(task))
	}

	return vo.List(listUserResp, total), nil
}

func (s *UserSrv) DeleteUser(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	err = dao.NewUserDao(ctx).DeleteUser(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}
