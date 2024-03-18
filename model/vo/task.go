package vo

import "github.com/ncuhome/story-cook/model/dao"

type TaskResp struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
}

func BuildTaskResp(task *dao.Task) *TaskResp {
	return &TaskResp{
		ID:       task.ID,
		Title:    task.Title,
		Content:  task.Content,
		CreateAt: task.CreatedAt.Format("2006/01/02"),
	}
}
