package routers

import (
	"chat/controllers/accounts"
	"chat/controllers/msg"
	"chat/controllers/normal"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "chat/docs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func initFile() {
	gin.DisableConsoleColor()
	logFile := viper.GetString("app.logFile")
	f, err := os.Create(logFile)
	if err != nil {
		fmt.Printf("创建日志文件失败：%s\n",err.Error())
		return
	}
	gin.DefaultWriter = io.MultiWriter(f)
}

func Init() {
	initFile()
	router := gin.Default()
	user_ := router.Group("/user")
	{
		user_.GET("/conversations", msg.GetConversations)
		user_.POST("/register", accounts.Register)
		user_.POST("/login", accounts.Login)
		user_.GET("/record", msg.GetChatRecord)
		user_.POST("/avatar",accounts.UploadAvatar)
		user_.POST("/modify",accounts.ModifyInfo)
	}
	index := router.Group("/index")
	{
		index.GET("/neardynamic", normal.NearDynamic)
		index.POST("/dynamic", normal.PublishDynamic)
		index.POST("/loveintro",accounts.PublishLoveIntro)
	}
	resource := router.Group("/resource")
	{
		resource.StaticFS("/upload",http.Dir("D:\\GoWork\\images"))
		resource.POST("/uploadimg", normal.UploadImg)
	}
	// msg := router.Group("/msg")
	// {
	// 	//msg.POST("/sendTxtMsg")
	// }
	httpPort := viper.GetString("app.httpPort")
	fmt.Println("httpPort:", httpPort)
	router.GET("/api/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
