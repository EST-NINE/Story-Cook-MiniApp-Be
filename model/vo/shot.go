package vo

import "github.com/ncuhome/story-cook/model/dao"

type ShotResp struct {
	UserId      uint `json:"user_id"`
	DishId      uint `json:"dish_id"`
	DishAmount  int  `json:"dish_amount"`
	IsUnlock    bool `json:"is_unlock"`     // 是否解锁
	IsFirstShot bool `json:"is_first_shot"` // 是否是第一次解锁
	PieceAmount int  `json:"piece_amount"`  // 解锁后获得的碎片数量
}

func BuildShotResp(userDish *dao.UserDish, isFirstShot bool, pieceAmount int) *ShotResp {
	return &ShotResp{
		UserId:      userDish.UserId,
		DishId:      userDish.DishId,
		DishAmount:  userDish.DishAmount,
		IsUnlock:    userDish.IsUnlock,
		IsFirstShot: isFirstShot,
		PieceAmount: pieceAmount,
	}
}
