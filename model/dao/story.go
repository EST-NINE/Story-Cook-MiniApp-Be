package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Story struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uint      `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"type:longtext" json:"content"`
	Count     uint      `gorm:"default:0" json:"count"`
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

// ListAllStory 得到全部故事列表
func (dao *StoryDao) ListAllStory(page, limit int) (stories []*Story, total int64, err error) {
	err = dao.DB.Model(&Story{}).
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

// DeleteStory 删除故事
func (dao *StoryDao) DeleteStory(id uint) error {
	return dao.DB.Model(&Story{}).Where("id = ?", id).Delete(&Story{}).Error
}

// UpdateStory 更新故事
func (dao *StoryDao) UpdateStory(id uint, story *Story) error {
	return dao.DB.Model(&Story{}).Where("id = ?", id).Updates(story).Error
}
