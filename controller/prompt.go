package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/tongyi"
	"github.com/ncuhome/story-cook/pkg/util"
	"net/http"
)

func GetPromptHandler(ctx *gin.Context) {
	var data dto.PromptDto
	data.ExtendStory = tongyi.ExtendStoryChara
	data.AssessStory = tongyi.AssessStoryChara
	data.EndStory = tongyi.EndStoryChara

	resp := vo.SuccessWithData(data)
	ctx.JSON(http.StatusOK, resp)
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
