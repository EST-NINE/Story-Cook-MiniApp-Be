package vo

import "github.com/ncuhome/story-cook/model/dao"

type ShotResp struct {
	UserId      uint `json:"user_id"`
	DishId      uint `json:"dish_id"`
	DishAmount  uint `json:"dish_amount"`
	PieceAmount uint `json:"piece_amount"`
	IsFirstShot bool `json:"is_first_shot"`
}

func BuildShotResp(userDish *dao.UserDish, isFirstShot bool) *ShotResp {
	return &ShotResp{
		UserId:      userDish.UserId,
		DishId:      userDish.DishId,
		DishAmount:  userDish.DishAmount,
		PieceAmount: userDish.PieceAmount,
		IsFirstShot: isFirstShot,
	}
}
