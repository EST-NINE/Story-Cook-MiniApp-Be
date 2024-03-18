package dao

import (
	"context"

	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Image       string `gorm:"not null"`
}

type DishDao struct {
	*gorm.DB
}

func NewDishDao(c context.Context) *DishDao {
	if c == nil {
		c = context.Background()
	}
	return &DishDao{NewDBClient(c)}
}

func (dao *DishDao) CreateDish(dish *Dish) error {
	return dao.DB.Model(&Dish{}).Create(&dish).Error
}

func (dao *DishDao) FindDishById(id uint) (dish *Dish, err error) {
	err = dao.DB.Model(&Dish{}).Where("id = ?", id).First(&dish).Error
	return dish, err
}

func (dao *DishDao) DeleteDish(id uint) error {
	return dao.DB.Model(&Dish{}).Where("id = ?", id).Delete(&Dish{}).Error
}

func (dao *DishDao) UpdateDish(id uint, dish *Dish) error {
	return dao.DB.Model(&Dish{}).Where("id = ?", id).Updates(dish).Error
}

func (dao *DishDao) ListDish() (dishes []*Dish, total int64, err error) {
	return dishes, total, dao.DB.Model(&Dish{}).Count(&total).Find(&dishes).Error
}
