package vo

import "github.com/ncuhome/story-cook/model/dao"

type StoryResp struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Mood      string `json:"mood"`
	Keywords  string `json:"keywords"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func BuildStoryResp(story *dao.Story) *StoryResp {
	return &StoryResp{
		ID:        story.ID,
		Title:     story.Title,
		Mood:      story.Mood,
		Keywords:  story.Keywords,
		Content:   story.Content,
		CreatedAt: story.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
