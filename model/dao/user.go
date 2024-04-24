package dao

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Openid   string `gorm:"primaryKey"`
	Money    int    `gorm:"default:0"`
	Piece    int    `gorm:"default:0"`
}

type UserDao struct {
	*gorm.DB
}

func NewUserDao(c context.Context) *UserDao {
	if c == nil {
		c = context.Background()
	}
	return &UserDao{NewDBClient(c)}
}

func (dao *UserDao) CreateUser(user *User) error {
	return dao.DB.Model(&User{}).Create(user).Error
}

func (dao *UserDao) FindUserByUserId(id uint) (user *User, err error) {
	err = dao.DB.Model(&User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (dao *UserDao) FindUserByOpenid(openid string) (user *User, err error) {
	err = dao.DB.Model(&User{}).Where("openid = ?", openid).First(&user).Error
	return user, err
}

func (dao *UserDao) UpdateUserById(id uint, user *User) error {
	return dao.DB.Model(&User{}).Where("id = ?", id).Updates(user).Error
}

func (dao *UserDao) DailyLoginReward(user *User) error {
	return dao.DB.Model(&User{}).Where("id = ?", user.ID).Update("money", gorm.Expr("money + ?", 20)).Error
}

func (dao *UserDao) ListUserByID(page int, limit int) (users []*User, total int64, err error) {
	err = dao.DB.Model(&User{}).
		Count(&total).
		Order("id DESC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&users).Error
	return users, total, err
}

func (dao *UserDao) ListUserByMoney(page int, limit int) (users []*User, total int64, err error) {
	err = dao.DB.Model(&User{}).
		Count(&total).
		Order("money DESC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&users).Error
	return users, total, err
}

func (dao *UserDao) ListUserByPiece(page int, limit int) (users []*User, total int64, err error) {
	err = dao.DB.Model(&User{}).
		Count(&total).
		Order("piece DESC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&users).Error
	return users, total, err
}

func (dao *UserDao) DeleteUser(id uint) error {
	return dao.DB.Model(&User{}).Where("id = ?", id).Delete(&User{}).Error
}
