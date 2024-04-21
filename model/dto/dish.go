package dto

type DishDto struct {
	ID          uint   `json:"id" example:"1"`
	Name        string `json:"name" example:"dish1"`
	Description string `json:"description" example:"dish1"`
	Image       string `json:"image" example:"dish1.png"`
	Quality     string `json:"quality" example:"SSR"`
}
