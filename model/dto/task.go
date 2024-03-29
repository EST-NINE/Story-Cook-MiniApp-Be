package dto

type TaskDto struct {
	ID      uint   `json:"id" example:"1"`
	Title   string `json:"title" example:"task1"`
	Content string `json:"content" example:"content1"`
}
