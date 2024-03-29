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

type StorySrv struct {
}

// CreateStory 创建故事
func (s *StorySrv) CreateStory(ctx *gin.Context, req *dto.StoryDto) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	story := dao.Story{
		UserId:  userInfo.Id,
		Title:   req.Title,
		Content: req.Content,
		Count:   0,
	}

	err = dao.NewStoryDao(ctx).CreateStory(&story)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

func (s *StorySrv) FindStoryById(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	story, err := dao.NewStoryDao(ctx).FindStoryById(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorNotExistStory), err
	}

	return vo.SuccessWithData(vo.BuildStoryResp(story)), nil
}

// ListStory 得到对应用户的故事
func (s *StorySrv) ListStory(ctx *gin.Context, req *dto.ListDto) (resp *vo.Response, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	stories, total, err := dao.NewStoryDao(ctx).ListStory(req.Page, req.Limit, userInfo.Id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	listStoryResp := make([]*vo.StoryResp, 0)
	for _, story := range stories {
		listStoryResp = append(listStoryResp, vo.BuildStoryResp(story))
	}

	return vo.List(listStoryResp, total), nil
}

// DeleteStory 删除故事
func (s *StorySrv) DeleteStory(ctx *gin.Context, id uint) (resp *vo.Response, err error) {
	err = dao.NewStoryDao(ctx).DeleteStory(id)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

// UpdateStory 更新故事
func (s *StorySrv) UpdateStory(ctx *gin.Context, req *dto.StoryDto) (resp *vo.Response, err error) {
	storyDao := dao.NewStoryDao(ctx)
	story := &dao.Story{
		Title:   req.Title,
		Content: req.Content,
	}

	err = storyDao.UpdateStory(req.ID, story)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}

// 判断故事次数是否超过限制
func (s *StorySrv) checkStoryCount(ctx *gin.Context, storyId uint) (resp *vo.Response, err error) {
	story, err := dao.NewStoryDao(ctx).FindStoryById(storyId)
	if err != nil {
		return vo.Error(err, myErrors.ErrorNotExistStory), err
	}

	if story.Count >= 3 {
		err = errors.New("故事次数已用完")
		return vo.Error(err), err
	}

	return vo.Success(), nil
}

// 增加故事次数
func (s *StorySrv) addStoryCount(ctx *gin.Context, storyId uint) (resp *vo.Response, err error) {
	story, err := dao.NewStoryDao(ctx).FindStoryById(storyId)
	if err != nil {
		return vo.Error(err, myErrors.ErrorNotExistStory), err
	}

	story.Count += 1
	err = dao.NewStoryDao(ctx).UpdateStory(storyId, story)
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.Success(), nil
}
