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
	DishAmount  int    `json:"dish_amount"`
	IsUnlock    bool   `gorm:"is_unlock" json:"is_unlock"`
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
		Joins("LEFT JOIN user_dish ud ON ud.dish_id = dish.id AND ud.user_id = ?", userId).
		Where("dish.deleted_at IS NULL").
		Order("ud.is_unlock desc, ud.dish_amount desc").
		Scan(&userDishList).Error
	return userDishList, err
}

func (dao *UserDishDao) DeleteUserDishByDishId(dishId uint) error {
	return dao.DB.Model(&UserDish{}).Where("dish_id = ?", dishId).Delete(&UserDish{}).Error
}
