package vo

import (
	"github.com/ncuhome/story-cook/model/dao"
)

type UserResp struct {
	ID       uint   `json:"id"`        // 用户ID
	UserName string `json:"user_name"` // 用户名
	Money    int    `json:"money"`     // 货币
	Piece    int    `json:"piece"`     // 碎片
	CreateAt string `json:"create_at"` // 创建
}

type TokenDataResp struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

func BuildUserResp(user *dao.User) *UserResp {
	return &UserResp{
		ID:       user.ID,
		UserName: user.UserName,
		Money:    user.Money,
		Piece:    user.Piece,
		CreateAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func SuccessWithDataAndToken(user *dao.User, token string) *Response {
	respData := &TokenDataResp{
		User:  BuildUserResp(user),
		Token: token,
	}
	return SuccessWithData(respData)
}
