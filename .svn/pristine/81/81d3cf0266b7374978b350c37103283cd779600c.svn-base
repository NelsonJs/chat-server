package main

import (
	"fmt"
	"io"
	"os"

	"chat/db/redis_serve"
	"chat/routers"
	"chat/service/websocket"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var redisManager *redis_serve.RedisManager

func main() {
	initConfig()
	initFile()
	//routers.Init()
	routers.InitScocketRouters()
	redisManager = redis_serve.ConnectRedis()
	websocket.StartWebSocket(redisManager)
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
	fmt.Println("config app:", viper.Get("app"))
}
