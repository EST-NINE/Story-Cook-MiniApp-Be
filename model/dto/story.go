package dto

type StoryDto struct {
	ID      uint   `json:"id" example:"1"`
	Title   string `json:"title" example:"story2"`
	Content string `json:"content" example:"content2"`
}

type ListStoryDto struct {
	Page  int `json:"page" example:"1"`
	Limit int `json:"limit" example:"10"`
}

type ExtendStoryDto struct {
	Title      string `json:"title" example:"林黛玉倒拔垂杨柳"`
	Background string `json:"background" example:"略"`
	Keywords   string `json:"keywords" example:"贾宝玉 薛宝钗"`
}

type AssessStoryDto struct {
	Title   string `json:"title" example:"林黛玉倒拔垂杨柳"`
	Content string `json:"Content" example:"略"`
}
