package controller

import (
	"net/http"

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
