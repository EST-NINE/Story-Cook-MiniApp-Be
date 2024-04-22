package controller

import (
	"net/http"

	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/pkg/util"
	"github.com/ncuhome/story-cook/service"
)

func ShotSingleHandler(ctx *gin.Context) {
	shotSrv := service.ShotSrv{}
	resp, err := shotSrv.SingleShot(ctx)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func TenShotsHandler(ctx *gin.Context) {
	shotSrv := service.ShotSrv{}
	resp, err := shotSrv.TenShots(ctx)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func MergePieceHandler(ctx *gin.Context) {
	var req dto.PieceDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	shotSrv := service.ShotSrv{}
	resp, err := shotSrv.MergePiece(ctx, &req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
