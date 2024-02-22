package dao

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	AdminName      string `gorm:"unique"`
	PasswordDigest string
}

// SetPassword 设置密码
func (admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordDigest), []byte(password))
	return err == nil
}

type AdminDao struct {
	*gorm.DB
}

func NewAdminDao(c context.Context) *AdminDao {
	if c == nil {
		c = context.Background()
	}
	return &AdminDao{NewDBClient(c)}
}

func (dao *AdminDao) CreateAdmin(admin *Admin) error {
	return dao.DB.Model(&Admin{}).Create(admin).Error
}

func (dao *AdminDao) FindAdminByAdminId(id uint) (admin *Admin, err error) {
	err = dao.DB.Model(&Admin{}).Where("id = ?", id).First(&admin).Error
	return admin, err
}

func (dao *AdminDao) FindAdminByAdminName(name string) (admin *Admin, err error) {
	err = dao.DB.Model(&Admin{}).Where("admin_name = ?", name).First(&admin).Error
	return admin, err
}

func (dao *AdminDao) UpdateAdminById(id uint, admin *Admin) error {
	return dao.DB.Model(&Admin{}).Where("id = ?", id).Updates(admin).Error
}
