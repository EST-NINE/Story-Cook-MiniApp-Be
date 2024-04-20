package controller

import (
	"net/http"
	"strconv"

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
	idStr := ctx.Param("dishID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	shotSrv := service.ShotSrv{}
	resp, err := shotSrv.MergePiece(ctx, uint(id))
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
