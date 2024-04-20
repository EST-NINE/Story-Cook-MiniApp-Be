package dao

import (
	"context"

	"gorm.io/gorm"
)

type UserDish struct {
	UserId      uint `gorm:"user_id" json:"user_id"`
	DishId      uint `gorm:"dish_id" json:"dish_id"`
	DishAmount  uint `gorm:"dish_amount" json:"dish_amount"`
	PieceAmount uint `gorm:"piece_amount" json:"piece_amount"`
}

type UserDishDao struct {
	*gorm.DB
}

func NewUserDishDao(c context.Context) *UserDishDao {
	if c == nil {
		c = context.Background()
	}
	return &UserDishDao{NewDBClient(c)}
}

func (dao *UserDishDao) CreateUserDish(userDish *UserDish) error {
	return dao.DB.Model(&UserDish{}).Create(&userDish).Error
}

func (dao *UserDishDao) FindUserDish(userId, dishId uint) (userDish *UserDish, err error) {
	err = dao.DB.Model(&UserDish{}).Where("user_id = ? AND dish_id = ?", userId, dishId).First(&userDish).Error
	return userDish, err
}

func (dao *UserDishDao) UpdateUserDish(userDish *UserDish) error {
	return dao.DB.Model(&UserDish{}).Where("user_id = ? AND dish_id = ?", userDish.UserId, userDish.DishId).Save(&userDish).Error
}
