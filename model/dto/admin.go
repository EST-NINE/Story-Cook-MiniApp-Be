package dto

type AdminDto struct {
	AdminName string `json:"admin_name" example:"john"`
	Password  string `json:"password" example:"12345678"`
}
