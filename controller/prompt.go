package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/tongyi"
	"github.com/ncuhome/story-cook/pkg/util"
	"net/http"
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

func UpdatePromptHandler(ctx *gin.Context) {
	var req dto.PromptDto

	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	if req.ExtendStory != "" {
		tongyi.ExtendStoryChara = req.ExtendStory
	}
	if req.EndStory != "" {
		tongyi.EndStoryChara = req.EndStory
	}
	if req.AssessStory != "" {
		tongyi.AssessStoryChara = req.AssessStory
	}
	ctx.JSON(http.StatusOK, vo.Success())
}
