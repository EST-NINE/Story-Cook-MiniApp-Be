package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
	"github.com/ncuhome/story-cook/service"
)

func CreateUserTaskHandler(ctx *gin.Context) {
	var req dto.UserTaskDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	taskSrv := service.UserTaskSrv{}
	resp, err := taskSrv.CreateUserTask(ctx, &req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func GetUserTaskHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	taskSrv := service.UserTaskSrv{}
	resp, err := taskSrv.FindUserTaskById(ctx, uint(id))
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func DeleteUserTaskHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	taskSrv := service.UserTaskSrv{}
	resp, err := taskSrv.DeleteUserTask(ctx, uint(id))
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func UpdateUserTaskHandler(ctx *gin.Context) {
	var req dto.UserTaskDto
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	taskSrv := service.UserTaskSrv{}
	resp, err := taskSrv.UpdateUserTask(ctx, &req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func ListUserTaskHandler(ctx *gin.Context) {
	limitStr := ctx.Param("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, vo.Error(err, myErrors.ErrorInvalidParams))
		return
	}

	taskSrv := service.UserTaskSrv{}
	resp, err := taskSrv.ListUserTask(ctx, limit)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
