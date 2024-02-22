package controller

import (
	"net/http"

	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/pkg/util"
	"github.com/ncuhome/story-cook/service"
)

func AdminRegisterHandler(ctx *gin.Context) {
	var req dto.AdminDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	adminSrv := service.AdminSrv{}
	resp, err := adminSrv.Register(ctx, &req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func AdminLoginHandler(ctx *gin.Context) {
	var req dto.AdminDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	adminSrv := service.AdminSrv{}
	resp, err := adminSrv.Login(ctx, &req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func GetAdminInfoHandler(ctx *gin.Context) {
	adminSrv := service.AdminSrv{}
	resp, err := adminSrv.AdminInfo(ctx)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func UpdateAdminInfoHandler(ctx *gin.Context) {
	var req dto.AdminDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err))
		return
	}

	adminSrv := service.AdminSrv{}
	resp, err := adminSrv.UpdateInfo(ctx, &req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
