package main

import (
	"blogbk/common"
	"blogbk/model"
	"blogbk/router"
	"flag"
)

func main() {
	f := flag.String("config", "./configs/config.toml", "配置文件路径")
	common.InitConfig(f)
	common.InitMysql()
	dBMigrate()
	router.InitRouter()
	if err := router.Router.Run(common.Config.Listen); err != nil {
		common.Log.Fatalf("服务运行失败：%s", err.Error())
	}
}

func dBMigrate() {
	common.Mysql.AutoMigrate(&model.SUser{}, &model.SCategory{}, &model.STag{}, &model.SPost{})
}
