package dao

import (
	"context"
	"errors"

	"github.com/ncuhome/story-cook/model/dto"
	"github.com/ncuhome/story-cook/pkg/tongyi"
	"github.com/ncuhome/story-cook/pkg/util"
	"gorm.io/gorm"
)

type Prompt struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	Type    string `gorm:"not null" json:"type"`
	Content string `gorm:"type:longtext" json:"content"`
}

type PromptList struct {
	ExtendStoryPrompt Prompt `json:"extend_story_prompt"`
	EndStoryPrompt    Prompt `json:"end_story_prompt"`
	AssessStoryPrompt Prompt `json:"assess_story_prompt"`
}

type PromptDao struct {
	*gorm.DB
}

func NewPromptDao(c context.Context) *PromptDao {
	if c == nil {
		c = context.Background()
	}
	return &PromptDao{NewDBClient(c)}
}

func (dao *PromptDao) FindPrompt() (PromptList, error) {
	var promptList PromptList
	err := dao.Model(&Prompt{}).Where("type = ?", "extend").Last(&promptList.ExtendStoryPrompt).Error
	if err != nil {
		return promptList, err
	}

	err = dao.Model(&Prompt{}).Where("type = ?", "end").Last(&promptList.EndStoryPrompt).Error
	if err != nil {
		return promptList, err
	}

	err = dao.Model(&Prompt{}).Where("type = ?", "assess").Last(&promptList.AssessStoryPrompt).Error
	if err != nil {
		return promptList, err
	}

	return promptList, nil
}

func (dao *PromptDao) GetPrompt(data *dto.PromptDto) error {
	var prompt Prompt
	err := dao.Model(&Prompt{}).Where("type = ?", "extend").Last(&prompt).Error
	if err != nil {
		return err
	}
	data.ExtendStory = prompt.Content

	prompt = Prompt{}
	err = dao.Model(&Prompt{}).Where("type = ?", "end").Last(&prompt).Error
	if err != nil {
		return err
	}
	data.EndStory = prompt.Content

	prompt = Prompt{}
	err = dao.Model(&Prompt{}).Where("type = ?", "assess").Last(&prompt).Error
	if err != nil {
		return err
	}
	data.AssessStory = prompt.Content
	return nil
}

func (dao *PromptDao) UpdatePrompt(req *dto.PromptDto) error {
	var prompt Prompt

	if req.ExtendStory != "" {
		err := dao.Model(&Prompt{}).Where("type = ?", "extend").Last(&prompt).Error
		if err != nil {
			return err
		}
		err = dao.Model(&prompt).Updates(&Prompt{Content: req.ExtendStory}).Error
		if err != nil {
			return err
		}
		tongyi.ExtendStoryPrompt = req.ExtendStory
	}

	if req.EndStory != "" {
		prompt = Prompt{}
		err := dao.Model(&Prompt{}).Where("type = ?", "end").Last(&prompt).Error
		if err != nil {
			return err
		}
		err = dao.Model(&prompt).Updates(&Prompt{Content: req.EndStory}).Error
		if err != nil {
			return err
		}
		tongyi.EndStoryPrompt = req.EndStory
	}

	if req.AssessStory != "" {
		prompt = Prompt{}
		err := dao.Model(&Prompt{}).Where("type = ?", "assess").Last(&prompt).Error
		if err != nil {
			return err
		}
		err = dao.Model(&prompt).Updates(&Prompt{Content: req.AssessStory}).Error
		if err != nil {
			return err
		}
		tongyi.AssessStoryPrompt = req.AssessStory
	}
	return nil
}

func InitPrompt() {
	var data dto.PromptDto

	promptDao := NewPromptDao(context.TODO())

	// 判断 prompt 表是否为空, 为空则新建三条记录
	result := promptDao.First(&Prompt{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		prompts := []*Prompt{
			{Type: "extend"},
			{Type: "end"},
			{Type: "assess"},
		}
		promptDao.Create(prompts)
	}

	err := promptDao.GetPrompt(&data)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	tongyi.ExtendStoryPrompt = data.ExtendStory
	tongyi.EndStoryPrompt = data.EndStory
	tongyi.AssessStoryPrompt = data.AssessStory
}
