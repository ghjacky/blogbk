package common

import "github.com/spf13/viper"

type SConfig struct {
	Listen string
	Log    string
	Salt   string
	SMysqlConfig
}

var Config = new(SConfig)

func InitConfig(file *string) {
	viper.SetConfigFile(*file)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		Log.Fatalf("配置文件读取错误：%s", err.Error())
	}
	Config.Listen = viper.GetString("main.listen")
	Config.Log = viper.GetString("main.log")
	Config.Salt = viper.GetString("main.salt")
	Config.SMysqlConfig.Host = viper.GetString("mysql.host")
	Config.SMysqlConfig.User = viper.GetString("mysql.user")
	Config.SMysqlConfig.Pass = viper.GetString("mysql.pass")
	Config.SMysqlConfig.Db = viper.GetString("mysql.db")
	Config.SMysqlConfig.ConnParams = "charset=utf8&parseTime=True&loc=Local"
}
