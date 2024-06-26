package dao

import (
	"context"
	"github.com/ncuhome/story-cook/pkg/global"
	"github.com/ncuhome/story-cook/pkg/util"

	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Image       string `gorm:"not null"`
	Quality     string `gorm:"not null"`
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

func (dao *DishDao) ListDishByQuality(quality string) (dishes []*Dish, err error) {
	return dishes, dao.DB.Model(&Dish{}).Where("quality = ?", quality).Find(&dishes).Error
}

func (dao *DishDao) ListDishesByQualities(qualities []string) (map[string][]*Dish, error) {
	var dishes []*Dish
	dishesMap := make(map[string][]*Dish)

	err := dao.DB.Model(&Dish{}).Where("quality IN (?)", qualities).Find(&dishes).Error
	if err != nil {
		return nil, err
	}

	for _, dish := range dishes {
		dishesMap[dish.Quality] = append(dishesMap[dish.Quality], dish)
	}

	return dishesMap, nil
}

func InitDishMap() {
	global.DishMap = make(map[string]string)
	dishes, total, err := NewDishDao(context.TODO()).ListDish()
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	var i int64
	for i = 0; i < total; i++ {
		global.DishMap[dishes[i].Name] = dishes[i].Description
	}
}
