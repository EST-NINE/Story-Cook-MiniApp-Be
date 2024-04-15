package config

import (
	"log"

	"gopkg.in/ini.v1"
)

func InitDevFile() {
	file, err := ini.Load("./config/config_dev.ini")
	if err != nil {
		log.Print(err)
		panic(err)
	}
	LoadServer(file)
	LoadMysqlData(file)
	LoadTongYi(file)
	LoadWx(file)
	LoadAliOss(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadTongYi(file *ini.File) {
	ApiKey = file.Section("tong_yi").Key("ApiKey").String()
}

func LoadWx(file *ini.File) {
	WxAppId = file.Section("wx").Key("WxAppId").String()
	WxAppSecret = file.Section("wx").Key("WxAppSecret").String()
}

func LoadAliOss(file *ini.File) {
	AliEndPoint = file.Section("AliOss").Key("AliEndPoint").String()
	AliAccessKeyId = file.Section("AliOss").Key("AliAccessKeyId").String()
	AliAccessKeySecret = file.Section("AliOss").Key("AliAccessKeySecret").String()
	AliBucketName = file.Section("AliOss").Key("AliBucketName").String()
}
