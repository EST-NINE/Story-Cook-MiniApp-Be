package main

import (
	"log"

	"github.com/ncuhome/story-cook/config"
	"github.com/ncuhome/story-cook/model/dao"
	"github.com/ncuhome/story-cook/pkg/util"
	"github.com/ncuhome/story-cook/router"
)

func main() {
	r := router.NewRouter()
	err := r.Run(config.HttpPort)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	// config.InitDevFile()
	util.InitLog()
	dao.InitMysql()
}
