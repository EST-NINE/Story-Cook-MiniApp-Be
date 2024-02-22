package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ncuhome/story-cook/config"
)

const grantType = "authorization_code"

type WechatLoginResponse struct {
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
}

func GetWxOpenid(wxCode string) (openid string, err error) {
	// 调用微信登录接口获取 session_key 和 openid
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=%s",
		config.WxAppId,
		config.WxAppSecret,
		wxCode,
		grantType,
	)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("wxResp:", string(body))

	var loginResp WechatLoginResponse
	if err := json.Unmarshal(body, &loginResp); err != nil {
		return "", err
	}

	return loginResp.OpenID, nil
}
