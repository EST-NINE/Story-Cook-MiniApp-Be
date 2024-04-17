package dao

import (
	"context"
	"github.com/ncuhome/story-cook/model/dto"
	"gorm.io/gorm"
)

type Prompt struct {
	ID      uint   `gorm:"primarykey"`
	Type    string `gorm:"not null"`
	Content string `gorm:"type:longtext"`
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
	}
	return nil
}
