package controller

import (
	"net/http"

	"github.com/ncuhome/story-cook/model/vo"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ncuhome/story-cook/pkg/util"
)

func UploadImageHandler(ctx *gin.Context) {
	// 获取前端传递的图片
	file, err := ctx.FormFile("file")
	if err != nil {
		return
	}

	// 拼接uuid的图片名称
	imageName := uuid.New().String() + file.Filename
	imagePath, err := util.AliOss(imageName, file)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, vo.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, vo.SuccessWithData(imagePath))
}
