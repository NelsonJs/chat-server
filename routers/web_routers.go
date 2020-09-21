package routers

import (
	"chat/config"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "chat/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func initFile() {
	gin.DisableConsoleColor()
	logFile := config.GetViperString("logFile")
	f, err := os.Create(logFile)
	if err != nil {
		fmt.Printf("创建日志文件失败：%s\n", err.Error())
		return
	}
	gin.DefaultWriter = io.MultiWriter(f)
}

func Init() {
	initFile()
	router := gin.Default()

	resource := router.Group("/resource")
	{
		resource.StaticFS("/upload", http.Dir("/dist/images"))

	}
	httpPort := config.GetViperString("httpPort")
	fmt.Println("httpPort:", httpPort)
	router.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
