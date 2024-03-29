package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ncuhome/story-cook/pkg/tongyi"

	"github.com/gin-gonic/gin"

	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
	"github.com/ncuhome/story-cook/service"
)

func CreateStoryHandler(ctx *gin.Context) {
	var req dto.StoryDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	storySrv := service.StorySrv{}
	resp, err := storySrv.CreateStory(ctx, &req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func ExtendStoryHandler(ctx *gin.Context) {
	var req dto.ExtendStoryDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	charaSetting := tongyi.ExtendStoryChara
	prompt := fmt.Sprintf("标题：%s 故事背景：%s 关键词：%s", req.Title, req.Background, req.Keywords)
	if err := ForWardSSE(ctx, prompt, charaSetting); err != nil {
		util.LogrusObj.Infoln(err)
		return
	}
}

func EndStoryHandler(ctx *gin.Context) {
	var req dto.ExtendStoryDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	charaSetting := tongyi.EndStoryChara
	prompt := fmt.Sprintf("标题：%s 故事背景：%s 关键词：%s", req.Title, req.Background, req.Keywords)
	if err := ForWardSSE(ctx, prompt, charaSetting); err != nil {
		util.LogrusObj.Infoln(err)
		return
	}
}

func AssessStoryHandler(ctx *gin.Context) {
	var req dto.AssessStoryDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	charaSetting := tongyi.AssessStoryChara
	prompt := fmt.Sprintf("故事标题：%s 故事内容：%s", req.Title, req.Content)
	if err := ForWardSSE(ctx, prompt, charaSetting); err != nil {
		util.LogrusObj.Infoln(err)
		return
	}
}

func GetStoryHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	storySrv := service.StorySrv{}
	resp, err := storySrv.FindStoryById(ctx, uint(id))
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func ListStoryHandler(ctx *gin.Context) {
	var req dto.ListDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	storySrv := service.StorySrv{}
	resp, err := storySrv.ListStory(ctx, &req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func DeleteStoryHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	storySrv := service.StorySrv{}
	resp, err := storySrv.DeleteStory(ctx, uint(id))
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func UpdateStoryHandler(ctx *gin.Context) {
	var req dto.StoryDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	storySrv := service.StorySrv{}
	resp, err := storySrv.UpdateStory(ctx, &req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
