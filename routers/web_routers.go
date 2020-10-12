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

func ListenRoute() {
	initFile()
	router := gin.Default()

	conversation := router.Group("/conversation")
	{
		conversation.GET("/list", ConversationList)
		conversation.GET("/record", ChatRecords)
		conversation.POST("/delete", DelConversation)
		conversation.POST("/revokeMsg", RevokeMsg)
	}

	group := router.Group("/group")
	{
		group.POST("/create", CreateGroup)
		group.POST("/updatename", UpdateName)
		group.POST("/addmanager", AddManager)
		group.POST("/addmember", AddMember)
		group.POST("/delmember", RemoveMember)
		group.POST("/updateavater", AddAvatar)
		group.POST("/transfer", Transfer) //转移群组
		group.POST("/del", Del)           //解散群组
		group.POST("/join", Join)
		group.POST("/exit", Exit)
	}

	feedBack := router.Group("/help")
	{
		feedBack.POST("/inform") //举报
	}

	resource := router.Group("/resource")
	{
		resource.StaticFS("/upload", http.Dir("/dist/images"))

	}
	httpPort := config.GetViperString("httpPort")
	fmt.Println("httpPort:", httpPort)
	router.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":5874")
}
