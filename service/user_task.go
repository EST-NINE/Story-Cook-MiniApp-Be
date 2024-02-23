package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
)

type UserTaskSrv struct {
}

func (s *UserTaskSrv) CreateUserTask(ctx *gin.Context, req *dto.UserTaskDto) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	task := dao.UserTask{
		UserId: userInfo.Id,
		TaskId: req.TaskId,
	}

	err = dao.NewUserTaskDao(ctx).CreateUserTask(&task)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

func (s *UserTaskSrv) FindUserTaskById(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	task, err := dao.NewUserTaskDao(ctx).GetUserTaskById(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorNotExistTask), err
	} else if task == nil {
		err = errors.New("task not found")
		return vo.Error(err, myErrors.ErrorNotExistTask), err
	}

	return vo.SuccessWithData(task), nil
}

func (s *UserTaskSrv) DeleteUserTask(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	err = dao.NewUserTaskDao(ctx).DeleteUserTask(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

func (s *UserTaskSrv) UpdateUserTask(ctx *gin.Context, req *dto.UserTaskDto) (resp *vo.Response, err error) {
	taskDao := dao.NewUserTaskDao(ctx)
	task, err := taskDao.FindUserTaskById(req.ID)
	if err != nil {
		return vo.Error(err, myErrors.ErrorNotExistTask), err
	}

	err = taskDao.UpdateUserTask(req.ID, task)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(task), nil
}

func (s *UserTaskSrv) ListUserTask(ctx *gin.Context, limit int) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	tasks, err := dao.NewUserTaskDao(ctx).ListUserTask(userInfo.Id, limit)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(tasks), nil
}