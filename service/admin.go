package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
	"gorm.io/gorm"
)

type AdminSrv struct {
}

func (s *AdminSrv) Register(ctx *gin.Context, req *dto.AdminDto) (resp *vo.Response, err error) {
	adminDao := dao.NewAdminDao(ctx)
	admin := &dao.Admin{
		AdminName: req.AdminName,
	}
	// 密码加密存储
	_ = admin.SetPassword(req.Password)

	if err = adminDao.CreateAdmin(admin); err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	token, err := util.GenerateToken(admin.ID, 1)
	if err != nil {
		return vo.Error(err, myErrors.ErrorAuthToken), err
	}

	respData := vo.AdminTokenDataResp{
		Admin: vo.BuildAdminResp(admin),
		Token: token,
	}
	return vo.SuccessWithData(respData), nil
}

func (s *AdminSrv) Login(ctx *gin.Context, req *dto.AdminDto) (resp *vo.Response, err error) {
	adminDao := dao.NewAdminDao(ctx)
	admin, err := adminDao.FindAdminByAdminName(req.AdminName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return vo.Error(err, myErrors.ErrorNotExistAdmin), err
	}

	if !admin.CheckPassword(req.Password) {
		err = errors.New("账号/密码错误")
		return vo.Error(err), err
	}

	token, err := util.GenerateToken(admin.ID, 1)
	if err != nil {
		return vo.Error(err), err
	}

	respData := vo.AdminTokenDataResp{
		Admin: vo.BuildAdminResp(admin),
		Token: token,
	}
	return vo.SuccessWithData(respData), nil
}

func (s *AdminSrv) AdminInfo(ctx *gin.Context) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	adminInfo := claims.(*util.Claims)

	adminDao := dao.NewAdminDao(ctx)
	admin, err := adminDao.FindAdminByAdminId(adminInfo.Id)
	if err != nil {
		return vo.Error(err), err
	}

	respData := vo.BuildAdminResp(admin)
	return vo.SuccessWithData(respData), nil
}

func (s *AdminSrv) UpdateInfo(ctx *gin.Context, req *dto.AdminDto) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	adminInfo := claims.(*util.Claims)

	adminDao := dao.NewAdminDao(ctx)
	admin := &dao.Admin{
		AdminName: req.AdminName,
	}

	err = adminDao.UpdateAdminById(adminInfo.Id, admin)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}
