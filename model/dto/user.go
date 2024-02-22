package dto

type UserDto struct {
	ID       uint   `json:"id" example:"1"`
	UserName string `json:"user_name" example:"john"`
	// 微信登录凭证 code
	Code string `json:"code" binding:"omitempty,len=32" example:"0a3FOa1w3Ek5c23Ey72w3l4HW02FOa1k"`
}
