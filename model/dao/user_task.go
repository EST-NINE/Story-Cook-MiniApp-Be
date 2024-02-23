package dao

import (
	"context"

	"gorm.io/gorm"
)

type UserTask struct {
	gorm.Model
	UserId  uint `gorm:"not null"`
	TaskId  uint `gorm:"not null"`
	StoryId uint `gorm:"not null"`
}

type UserTaskVO struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}

type UserTaskDao struct {
	*gorm.DB
}

func NewUserTaskDao(c context.Context) *UserTaskDao {
	if c == nil {
		c = context.Background()
	}
	return &UserTaskDao{NewDBClient(c)}
}

func (dao *UserTaskDao) CreateUserTask(task *UserTask) error {
	return dao.DB.Model(&UserTask{}).Create(&task).Error
}

func (dao *UserTaskDao) FindUserTaskById(id uint) (task *UserTask, err error) {
	err = dao.DB.Model(&UserTask{}).Where("id = ?", id).First(&task).Error
	return task, err
}

func (dao *UserTaskDao) DeleteUserTask(id uint) error {
	return dao.DB.Model(&UserTask{}).Where("id = ?", id).Delete(&UserTask{}).Error
}

func (dao *UserTaskDao) UpdateUserTask(id uint, task *UserTask) error {
	return dao.DB.Model(&UserTask{}).Where("id = ?", id).Updates(task).Error
}

func (dao *UserTaskDao) GetUserTaskById(id uint) (task *UserTaskVO, err error) {
	err = dao.DB.Table("task t").
		Select("t.id, t.content, ut.status").
		Joins("LEFT JOIN user_task ut ON t.id = ut.task_id").
		Where("ut.id = ? AND ut.deleted_at IS NULL", id).
		Scan(&task).
		Error
	return task, err
}

func (dao *UserTaskDao) ListUserTask(userId uint, limit int) (tasks []*UserTaskVO, err error) {
	err = dao.DB.Table("task t").
		Select("t.id, t.content, ut.status").
		Joins("LEFT JOIN user_task ut ON t.id = ut.task_id AND ut.user_id = ?", userId).
		Where("t.deleted_at IS NULL").
		Order("t.created_at DESC").
		Limit(limit).
		Scan(&tasks).
		Error
	return tasks, err
}
