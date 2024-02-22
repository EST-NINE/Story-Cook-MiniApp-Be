package config

import (
	"os"
)

var (
	AppMode  string = "release"
	HttpPort string = ":8082"

	Db         string = "mysql"
	DbHost     string = os.Getenv("DB_HOST")
	DbPort     string = os.Getenv("DB_PORT")
	DbUser     string = os.Getenv("DB_USER")
	DbPassWord string = os.Getenv("DB_PASSWORD")
	DbName     string = os.Getenv("DB_NAME")

	AppId     string = os.Getenv("APP_ID")
	ApiKey    string = os.Getenv("API_KEY")
	ApiSecret string = os.Getenv("API_SECRET")

	WxAppId     string = os.Getenv("WX_APP_ID")
	WxAppSecret string = os.Getenv("WX_APP_SECRET")
)
