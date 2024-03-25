package dao

import (
	"context"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"type:longtext"`
}

type UserTask struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
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

func (dao *TaskDao) DeleteTask(id uint) error {
	return dao.DB.Model(&Task{}).Where("id = ?", id).Delete(&Task{}).Error
}

func (dao *TaskDao) UpdateTask(id uint, task *Task) error {
	return dao.DB.Model(&Task{}).Where("id = ?", id).Updates(task).Error
}

func (dao *TaskDao) GetDailyTask(userId uint) (task *UserTask, err error) {
	err = dao.DB.Table("task").
		Select("task.id, task.title, task.content, IFNULL(orders.status, 0) as status").
		Joins("LEFT JOIN orders ON task.id = orders.task_id AND orders.user_id = ?", userId).
		Where("task.deleted_at IS NULL").
		Order("task.created_at DESC").
		Limit(1).
		Scan(&task).Error
	return task, err
}

func (dao *TaskDao) ListUserTask(userId uint, page int, limit int) (tasks []*UserTask, total int64, err error) {
	err = dao.DB.Table("task").
		Select("task.id, task.title, task.content, IFNULL(orders.status, 0) as status").
		Joins("LEFT JOIN orders ON task.id = orders.task_id AND orders.user_id = ?", userId).
		Where("task.deleted_at IS NULL").
		Order("task.created_at DESC").
		Offset((page - 1) * limit).
		Limit(limit).
		Scan(&tasks).Error
	return tasks, total, err
}
