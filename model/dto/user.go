package dto

type UserDto struct {
	UserID   uint   `json:"user_id" example:"1"`
	UserName string `json:"user_name" example:"john"`
	Code     string `json:"code" binding:"omitempty,len=32" example:"0a3FOa1w3Ek5c23Ey72w3l4HW02FOa1k"` // 微信登录凭证 code
	Money    int    `json:"money" binding:"omitempty" example:"100"`
	Piece    int    `json:"piece" binding:"omitempty" example:"10"`
}

type ListUserDto struct {
	Page  int  `json:"page" example:"1"`
	Limit int  `json:"limit" example:"10"`
	Order uint `json:"order" example:"0"` // 0: ID 1: 金钱数量
}
