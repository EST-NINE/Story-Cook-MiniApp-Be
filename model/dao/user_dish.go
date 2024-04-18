package dao

type UserDish struct {
	userId      uint `gorm:"user_id"`
	dishId      uint `gorm:"dish_id"`
	dishAmount  uint `gorm:"dish_amount"`
	pieceAmount uint `gorm:"piece_amount"`
}
