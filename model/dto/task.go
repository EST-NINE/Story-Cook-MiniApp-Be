package dto

type TaskDto struct {
	ID      uint   `json:"id" example:"1"`
	Content string `json:"content" binding:"required" example:"content1"`
}

type ListTaskDto struct {
	Page  int `json:"page" example:"1"`
	Limit int `json:"limit" example:"10"`
}
