package dao

import (
	"context"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	UserId  uint   `gorm:"not null"`
	TaskId  uint   `gorm:"not null"`
	StoryId uint   `gorm:"not null"`
	Comment string `gorm:"type:longtext"`
	Score   int    `gorm:"default:0"`
	Money   int    `gorm:"default:0"`
	Status  int    `gorm:"default:0"` // 0:未完成 1:进行中 2:已完成
}

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(c context.Context) *OrderDao {
	if c == nil {
		c = context.Background()
	}
	return &OrderDao{NewDBClient(c)}
}

func (dao *OrderDao) CreateOrder(order *Orders) error {
	return dao.DB.Model(&Orders{}).Create(&order).Error
}

func (dao *OrderDao) FindOrderById(id uint) (order *Orders, err error) {
	err = dao.DB.Model(&Orders{}).Where("id = ?", id).First(&order).Error
	return order, err
}

func (dao *OrderDao) DeleteOrder(id uint) error {
	return dao.DB.Model(&Orders{}).Where("id = ?", id).Delete(&Orders{}).Error
}

func (dao *OrderDao) UpdateOrder(id uint, order *Orders) error {
	return dao.DB.Model(&Orders{}).Where("id = ?", id).Updates(order).Error
}

func (dao *OrderDao) ListOrder(page, limit int, userId uint) (orders []*Orders, total int64, err error) {
	err = dao.DB.Model(&Orders{}).Where("user_id = ?", userId).
		Count(&total).
		Order("created_at DESC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&orders).Error
	return orders, total, err
}
