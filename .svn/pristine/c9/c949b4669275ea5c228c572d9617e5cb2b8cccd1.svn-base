package main

import (
	"fmt"
	"io"
	"os"

	"chat/db/mysql_serve"
	"chat/db/redis_serve"
	"chat/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var redisManager *redis_serve.RedisManager

func main() {
	initConfig()
	//initFile()
	routers.InitScocketRouters()
	mysql_serve.InitMySQL()
	//redisManager = redis_serve.ConnectRedis()
	//go websocket.StartWebSocket(redisManager)
	routers.Init()
}

func initFile() {
	gin.DisableConsoleColor()
	logFile := viper.GetString("app.logFile")
	f, _ := os.Create(logFile)
	gin.DefaultWriter = io.MultiWriter(f)
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
