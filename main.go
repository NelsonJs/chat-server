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

// @title 我的DEMO API
// @version 1.0     //当前api版本
// @description  swagger测试显示restful api  //说明信息
// @BasePath /user
// @contact.name youngxhu 名字
// @contact.url https://youngxhui.top 网址
// @contact.email youngxhui@g mail.com 邮箱
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080  //跟BasePath一样的显示效果
func main() {
	initConfig()
	routers.InitScocketRouters()
	mysql_serve.InitMySQL()
	//redisManager = redis_serve.ConnectRedis()
	go websocket.StartWebSocket(redisManager)
	routers.Init()
}



func initConfig() {
	viper.SetConfigName("config/app")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file:%s \n", err))
	}
	fmt.Println("config app:", viper.GetString("app"))
}
