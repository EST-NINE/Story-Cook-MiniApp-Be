package dto

type StoryDto struct {
	ID      uint   `json:"id" example:"1"`
	Title   string `json:"title" example:"story2"`
	Content string `json:"content" example:"content2"`
}

type ExtendStoryDto struct {
	StoryId  uint   `json:"storyId" example:"1"`
	Keywords string `json:"keywords" example:"贾宝玉 薛宝钗"`
}
