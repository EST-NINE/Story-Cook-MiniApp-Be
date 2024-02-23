package dto

type TaskDto struct {
	ID      uint   `json:"id" example:"1"`
	Title   string `json:"title" example:"task1"`
	Content string `json:"content" example:"content1"`
}

type ListTaskDto struct {
	Page  int `json:"page" example:"1"`
	Limit int `json:"limit" example:"10"`
}
