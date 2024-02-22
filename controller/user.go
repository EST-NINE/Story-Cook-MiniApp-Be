package controller

import (
	"net/http"

	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/service"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/pkg/util"
)

func UserLoginHandler(ctx *gin.Context) {
	var req dto.UserDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	userSrv := service.UserSrv{}
	resp, err := userSrv.Login(ctx, &req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func GetUserInfoHandler(ctx *gin.Context) {
	userSrv := service.UserSrv{}
	resp, err := userSrv.UserInfo(ctx)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func UpdateUserInfoHandler(ctx *gin.Context) {
	var req dto.UserDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err))
		return
	}

	userSrv := service.UserSrv{}
	resp, err := userSrv.UpdateInfo(ctx, &req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
