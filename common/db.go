package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type SMysqlConfig struct {
	Host       string
	User       string
	Pass       string
	Db         string
	ConnParams string
}

var Mysql *gorm.DB

func InitMysql() {
	var err error
	Log.Infoln("Connecting Mysql ......")
	Mysql, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		Config.SMysqlConfig.User, Config.SMysqlConfig.Pass, Config.SMysqlConfig.Host,
		Config.SMysqlConfig.Db, Config.SMysqlConfig.ConnParams))
	if err != nil {
		Log.Fatalf("Couldn't connect to mysql (%s): %s", Config.SMysqlConfig.Host, err.Error())
	} else {
		Log.Infof("Connected to mysql (%s) successfully", Config.SMysqlConfig.Host)
		Mysql = Mysql.LogMode(true)
	}
}
