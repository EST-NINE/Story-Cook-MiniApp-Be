package dto

type StoryDto struct {
	ID      uint   `json:"id" example:"1"`
	Title   string `json:"title" example:"story2"`
	Content string `json:"content" example:"content2"`
	Score   int    `json:"score" example:"100"`
	Money   int    `json:"money" example:"100"`
	Status  int    `json:"status" example:"1"`
}

type ListStoryDto struct {
	Page  int `json:"page" example:"1"`
	Limit int `json:"limit" example:"10"`
}
