package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
	"github.com/ncuhome/story-cook/service"
)

func GetPromptHandler(ctx *gin.Context) {
	var data dto.PromptDto

	err := dao.NewPromptDao(ctx).GetPrompt(&data)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, vo.Error(err, myErrors.ErrorDatabase))
		return
	}
	ctx.JSON(http.StatusOK, vo.SuccessWithData(data))
}

func FindPromptListHandler(ctx *gin.Context) {
	promptSrv := service.PromptSrv{}
	resp, err := promptSrv.FindPrompt(ctx)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func UpdatePromptHandler(ctx *gin.Context) {
	var req dto.PromptDto

	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	err := dao.NewPromptDao(ctx).UpdatePrompt(&req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, vo.Error(err, myErrors.ErrorDatabase))
		return
	}
	ctx.JSON(http.StatusOK, vo.Success())
}
