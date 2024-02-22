package dao

import (
	"context"

	"gorm.io/gorm"
)

type Story struct {
	gorm.Model
	UserId   uint   `gorm:"not null"`
	Title    string `gorm:"not null"`
	Mood     string `gorm:"default:开心"`
	Keywords string `gorm:"type:varchar(32)"`
	Content  string `gorm:"type:longtext"`
}

type StoryDao struct {
	*gorm.DB
}

func NewStoryDao(c context.Context) *StoryDao {
	if c == nil {
		c = context.Background()
	}
	return &StoryDao{NewDBClient(c)}
}

// CreateStory 创建故事
func (dao *StoryDao) CreateStory(story *Story) error {
	return dao.DB.Model(&Story{}).Create(&story).Error
}

// ListStory 得到故事列表
func (dao *StoryDao) ListStory(page, limit int, userId uint) (stories []*Story, total int64, err error) {
	err = dao.DB.Model(&Story{}).Where("user_id = ?", userId).
		Count(&total).
		Order("created_at DESC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&stories).Error
	return stories, total, err
}

// FindStoryById 根据故事id查找故事
func (dao *StoryDao) FindStoryById(id uint) (story *Story, err error) {
	err = dao.DB.Model(&Story{}).Where("id = ? ", id).First(&story).Error
	return story, err
}

// FindStoryByTitleAndUserId 根据故事title查找故事
func (dao *StoryDao) FindStoryByTitleAndUserId(uid uint, title string) (story *Story, err error) {
	err = dao.DB.Model(&Story{}).Where("user_id = ? AND title = ? ", uid, title).First(&story).Error
	return story, err
}

// DeleteStory 删除故事
func (dao *StoryDao) DeleteStory(id uint) error {
	story, err := dao.FindStoryById(id)
	if err != nil {
		return err
	}
	return dao.Delete(&story).Error
}

// UpdateStory 更新故事
func (dao *StoryDao) UpdateStory(id uint, story *Story) error {
	return dao.DB.Model(&Story{}).Where("id = ?", id).Updates(story).Error
}
