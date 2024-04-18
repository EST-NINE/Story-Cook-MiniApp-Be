package controller

import (
	"net/http"

	"github.com/ncuhome/story-cook/model/vo"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/pkg/util"
)

func UploadImageHandler(ctx *gin.Context) {
	// 获取前端传递的图片
	file, err := ctx.FormFile("file")
	if err != nil {
		return
	}

	imagePath, err := util.AliOss(file.Filename, file)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, vo.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, vo.SuccessWithData(imagePath))
}
