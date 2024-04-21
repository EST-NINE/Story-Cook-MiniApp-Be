package vo

import "github.com/ncuhome/story-cook/model/dao"

type DishResp struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Quality     string `json:"quality"`
	CreateAt    string `json:"create_at"`
}

func BuildDishResp(dish *dao.Dish) *DishResp {
	return &DishResp{
		ID:          dish.ID,
		Name:        dish.Name,
		Description: dish.Description,
		Image:       dish.Image,
		Quality:     dish.Quality,
		CreateAt:    dish.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
