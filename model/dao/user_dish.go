package dao

import (
	"context"

	"gorm.io/gorm"
)

type UserDish struct {
	UserId     uint `gorm:"user_id" json:"user_id"`
	DishId     uint `gorm:"dish_id" json:"dish_id"`
	DishAmount int  `gorm:"dish_amount" json:"dish_amount"`
	IsUnlock   bool `gorm:"is_unlock" json:"is_unlock"`
}

type UserDishResp struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Quality     string `json:"quality"`
	DishAmount  uint   `json:"dish_amount"`
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

func (dao *UserDishDao) ListUserDish(userId uint) (userDishList []*UserDishResp, err error) {
	err = dao.DB.Table("dish").
		Select("dish.id, dish.name, dish.description, dish.image, dish.quality, ud.dish_amount, ud.is_unlock").
		Joins("left join user_dish ud on ud.dish_id = dish.id").
		Where("ud.user_id = ?", userId).
		Scan(&userDishList).Error
	return userDishList, err
}
