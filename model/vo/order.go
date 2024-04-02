package vo

import "github.com/ncuhome/story-cook/model/dao"

type OrderResp struct {
	OrderId uint   `json:"order_id "`
	TaskId  uint   `json:"task_id "`
	StoryId uint   `json:"story_id "`
	Comment string `json:"comment "`
	Score   int    `json:"score "`
	Money   int    `json:"money "`
	Status  int    `json:"status "` // 0:未完成 1:进行中 2:已完成
}

func BuildOrderResp(order *dao.Orders) *OrderResp {
	return &OrderResp{
		OrderId: order.ID,
		TaskId:  order.TaskId,
		StoryId: order.StoryId,
		Comment: order.Comment,
		Score:   order.Score,
		Money:   order.Money,
		Status:  order.Status,
	}
}
