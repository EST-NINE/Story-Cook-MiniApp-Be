package vo

import "github.com/ncuhome/story-cook/model/dao"

type StoryResp struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Score     int    `json:"score"`
	Money     int    `json:"money"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
}

func BuildStoryResp(story *dao.Story) *StoryResp {
	return &StoryResp{
		ID:        story.ID,
		Title:     story.Title,
		Content:   story.Content,
		Score:     story.Score,
		Money:     story.Money,
		Status:    story.Status,
		CreatedAt: story.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
