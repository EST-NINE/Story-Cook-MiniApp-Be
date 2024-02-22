package dao

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Openid   string `gorm:"unique"`
	Money    int
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
