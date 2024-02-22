package dao

import (
	"context"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Content string `gorm:"not null"`
}

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(c context.Context) *TaskDao {
	if c == nil {
		c = context.Background()
	}
	return &TaskDao{NewDBClient(c)}
}

func (dao *TaskDao) CreateTask(task *Task) error {
	return dao.DB.Model(&Task{}).Create(&task).Error
}

func (dao *TaskDao) FindTaskById(id uint) (task *Task, err error) {
	err = dao.DB.Model(&Task{}).Where("id = ?", id).First(&task).Error
	return task, err
}

func (dao *TaskDao) ListTask(page, limit int) (tasks []*Task, total int64, err error) {
	err = dao.DB.Model(&Task{}).
		Order("created_at DESC").
		Count(&total).
		Offset((page - 1) * limit).
		Limit(limit).Find(&tasks).Error
	return tasks, total, err
}

func (dao *TaskDao) DeleteTask(id uint) error {
	story, err := dao.FindTaskById(id)
	if err != nil {
		return err
	}
	return dao.Delete(&story).Error
}

func (dao *TaskDao) UpdateTask(id uint, task *Task) error {
	return dao.DB.Model(&Task{}).Where("id = ?", id).Updates(task).Error
}
