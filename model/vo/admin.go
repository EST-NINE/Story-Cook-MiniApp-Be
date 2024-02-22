package vo

import "github.com/ncuhome/story-cook/model/dao"

type AdminResp struct {
	ID        uint   `json:"id"`
	AdminName string `json:"admin_name"`
	CreateAt  string `json:"create_at"`
}

type AdminTokenDataResp struct {
	Admin interface{} `json:"admin"`
	Token string      `json:"token"`
}

func BuildAdminResp(admin *dao.Admin) *AdminResp {
	return &AdminResp{
		ID:        admin.ID,
		AdminName: admin.AdminName,
		CreateAt:  admin.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func SuccessWithAdminAndToken(admin *dao.Admin, token string) *Response {
	respData := &AdminTokenDataResp{
		Admin: BuildAdminResp(admin),
		Token: token,
	}
	return SuccessWithData(respData)
}
