package dto

type CreateStoryDto struct {
	Title    string `json:"title" binding:"required"  example:"story1"`
	Mood     string `json:"mood" binding:"required" example:"开心"`
	Keywords string `json:"keywords" binding:"required" example:"室友+电脑"`
	Content  string `json:"content" binding:"required" example:"content1"`
}

type ListStoryDto struct {
	Page  int `json:"page" example:"1"`
	Limit int `json:"limit" example:"10"`
}

type DeleteStoryDto struct {
	ID uint `json:"id" binding:"required" example:"1"`
}

type UpdateStoryDto struct {
	ID            uint   `json:"id" binding:"required" example:"1"`
	UpdateTitle   string `json:"update_title" example:"story2"`
	UpdateContent string `json:"update_content" example:"content2"`
}
