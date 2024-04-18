package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/model/vo"
	"github.com/ncuhome/story-cook/pkg/myErrors"
)

type PromptSrv struct {
}

func (s *PromptSrv) FindPrompt(ctx *gin.Context) (resp *vo.Response, err error) {
	promptList, err := dao.NewPromptDao(ctx).FindPrompt()
	if err != nil {
		return vo.Error(err, myErrors.ErrorDatabase), err
	}

	return vo.SuccessWithData(promptList), nil
}
