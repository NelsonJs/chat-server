package main

import (
	"chat/db/mysql_serve"
	"chat/db/redis_serve"
	"chat/routers"
	"chat/service/websocket"
	"fmt"

	"github.com/spf13/viper"
)

var redisManager *redis_serve.RedisManager

// @title 微聊 API
// @version 1.0
// @description 不定时更新
// @BasePath
// @contact.name pjsong
// @contact.url xxx 网址
// @contact.email 18320944165@163.com 邮箱
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
func main() {
	initConfig()
	routers.InitScocketRouters()
	mysql_serve.InitMySQL()
	//redisManager = redis_serve.ConnectRedis()
	go websocket.StartWebSocket(redisManager)
	routers.Init()
}

func initConfig() {
	viper.SetConfigName("./config/app.yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file:%s \n", err))
	}
	fmt.Println("config app:", viper.GetString("app"))
}
