package dto

type UserTaskDto struct {
	ID     uint `json:"id" example:"1"`
	UserId uint `json:"user_id" example:"1"`
	TaskId uint `json:"task_id" example:"1"`
	Status int  `json:"status" example:"1"`
}

type ListUserTaskDto struct {
	UserId uint `json:"user_id" example:"1"`
	Limit  int  `json:"limit" example:"10"`
}
