package vo

import "github.com/ncuhome/story-cook/model/dao"

type TaskResp struct {
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
}

func BuildTaskResp(task *dao.Task) *TaskResp {
	return &TaskResp{
		ID:       task.ID,
		Content:  task.Content,
		CreateAt: task.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
