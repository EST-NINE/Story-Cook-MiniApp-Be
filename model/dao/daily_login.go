package dao

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type DailyLogin struct {
	UserId uint      `gorm:"primaryKey" json:"user_id"`
	Date   time.Time `json:"updated_at"`
}

type DailyLoginDao struct {
	*gorm.DB
}

func NewDailyLoginDao(c context.Context) *DailyLoginDao {
	if c == nil {
		c = context.Background()
	}
	return &DailyLoginDao{NewDBClient(c)}
}

func (dao *DailyLoginDao) CreateDailyLogin(dailyLogin *DailyLogin) error {
	return dao.DB.Model(&DailyLogin{}).Create(&dailyLogin).Error
}

func (dao *DailyLoginDao) FindDailyLoginById(userId uint) (dailyLogin *DailyLogin, err error) {
	today := time.Now()
	midnight := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	tomorrow := midnight.Add(24 * time.Hour)

	err = dao.DB.Model(&DailyLogin{}).Where("user_id = ? AND Date >= ? AND Date < ?", userId, midnight, tomorrow).First(&dailyLogin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_ = dao.DB.Model(&DailyLogin{}).Where("user_id = ?", userId).Update("Date", time.Now()).Error
	}

	return dailyLogin, err
}
