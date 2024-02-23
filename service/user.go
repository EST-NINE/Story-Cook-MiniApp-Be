package service

import (
	"errors"
	"fmt"

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
	user, err := userDao.FindUserByOpenid(openid)
	// 未找到，则创建一个用户
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &dao.User{
			UserName: fmt.Sprint("用户" + uuid.New().String()[:6]),
			Openid:   openid,
		}

		if err = userDao.CreateUser(user); err != nil {
			return vo.Error(err, myErrors.ErrorCreateUser), err
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return vo.Error(err, myErrors.ErrorNotExistUser), err
	}

	if req.UserName != "" {
		user.UserName = req.UserName
	}

	err = userDao.UpdateUserById(userInfo.Id, user)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}
