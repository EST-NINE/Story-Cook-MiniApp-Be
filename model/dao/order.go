package dao

import (
	"context"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId  uint   `gorm:"not null"`
	TaskId  uint   `gorm:"not null"`
	StoryId uint   `gorm:"not null"`
	Comment string `gorm:"type:longtext"`
	Score   int    `gorm:"default:0"`
	Money   int    `gorm:"default:0"`
	Status  int    `gorm:"default:0"` // 0:未完成 1:完成
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

func (dao *OrderDao) CreateOrder(order *Order) error {
	return dao.DB.Model(&Order{}).Create(&order).Error
}

func (dao *OrderDao) FindOrderById(id uint) (order *Order, err error) {
	err = dao.DB.Model(&Order{}).Where("id = ?", id).First(&order).Error
	return order, err
}

func (dao *OrderDao) DeleteOrder(id uint) error {
	return dao.DB.Model(&Order{}).Where("id = ?", id).Delete(&Order{}).Error
}

func (dao *OrderDao) UpdateOrder(id uint, order *Order) error {
	return dao.DB.Model(&Order{}).Where("id = ?", id).Updates(order).Error
}

func (dao *OrderDao) ListOrder(page, limit int, userId uint) (orders []*Order, total int64, err error) {
	err = dao.DB.Model(&Order{}).Where("user_id = ?", userId).
		Count(&total).
		Order("created_at DESC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&orders).Error
	return orders, total, err
}
