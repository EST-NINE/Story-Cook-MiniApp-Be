package dto

type OrderDto struct {
	ID      uint   `json:"id" example:"1"`
	UserId  uint   `json:"user_id" example:"1"`
	TaskId  uint   `json:"task_id" example:"1"`
	StoryId uint   `json:"story_id" example:"1"`
	Comment string `json:"comment" example:"100"`
	Score   int    `json:"score" example:"100"`
	Money   int    `json:"money" example:"100"`
	Status  int    `json:"status" example:"1"`
}
