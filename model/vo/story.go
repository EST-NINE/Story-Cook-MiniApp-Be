package vo

import "github.com/ncuhome/story-cook/model/dao"

type StoryResp struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Count     uint   `json:"count"`
	CreatedAt string `json:"created_at"`
}

func BuildStoryResp(story *dao.Story) *StoryResp {
	return &StoryResp{
		ID:        story.ID,
		Title:     story.Title,
		Content:   story.Content,
		Count:     story.Count,
		CreatedAt: story.CreatedAt.Format("2006/01/02"),
	}
}
