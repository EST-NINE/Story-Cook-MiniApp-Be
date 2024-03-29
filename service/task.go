package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
	"github.com/ncuhome/story-cook/pkg/util"
)

type TaskSrv struct {
}

func (s *TaskSrv) CreateTask(ctx *gin.Context, req *dto.TaskDto) (resp *vo.Response, err error) {
	task := dao.Task{
		Title:   req.Title,
		Content: req.Content,
	}

	err = dao.NewTaskDao(ctx).CreateTask(&task)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

func (s *TaskSrv) FindTaskById(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	task, err := dao.NewTaskDao(ctx).FindTaskById(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorNotExistTask), err
	}

	return vo.SuccessWithData(vo.BuildTaskResp(task)), nil
}

func (s *TaskSrv) DeleteTask(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	err = dao.NewTaskDao(ctx).DeleteTask(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

func (s *TaskSrv) UpdateTask(ctx *gin.Context, req *dto.TaskDto) (resp *vo.Response, err error) {
	taskDao := dao.NewTaskDao(ctx)
	task := &dao.Task{
		Title:   req.Title,
		Content: req.Content,
	}

	err = taskDao.UpdateTask(req.ID, task)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

func (s *TaskSrv) ListTask(ctx *gin.Context, req *dto.ListDto) (resp *vo.Response, err error) {
	tasks, total, err := dao.NewTaskDao(ctx).ListTask(req.Page, req.Limit)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	listTaskResp := make([]*vo.TaskResp, 0)
	for _, task := range tasks {
		listTaskResp = append(listTaskResp, vo.BuildTaskResp(task))
	}

	return vo.List(listTaskResp, total), nil
}

func (s *TaskSrv) GetDailyUserTask(ctx *gin.Context) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	task, err := dao.NewTaskDao(ctx).GetDailyTask(userInfo.Id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorNotExistTask), err
	}

	return vo.SuccessWithData(task), nil
}

func (s *TaskSrv) ListUserTask(ctx *gin.Context, req *dto.ListDto) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	tasks, total, err := dao.NewTaskDao(ctx).ListUserTask(userInfo.Id, req.Page, req.Limit)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.List(tasks, total), nil
}
