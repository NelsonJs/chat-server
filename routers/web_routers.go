package routers

import (
	"chat/business"
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

	test := router.Group("/test")
	{
		test.POST("/login",business.Test1)
		test.POST("/file",business.TestFile)
	}


	user := router.Group("/user")
	{
		user.POST("/register",business.Register)
		user.POST("/login",business.Login)
	}

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

	index := router.Group("/index")
	{
		index.GET("/neardynamic",business.NearDynamicList) //首页动态列表
		index.POST("/dynamic",business.InsertDynamic)
		index.POST("/dynamic/like",business.LikeDynamic)
	}

	travel := router.Group("/travel")
	{
		travel.GET("/list",business.GetTravel)
		travel.POST("/publish",business.PublishTravel)
		travel.POST("/join",business.JoinTravel)
		travel.POST("/exit",business.ExitTravel)
	}

	love := router.Group("/love")
	{
		love.GET("/list",business.GetLoveAll)
		love.PUT("/publish",business.PublishLove)
	}

	comments := router.Group("/comment")
	{
		comments.GET("/list",business.GetComments)
		comments.POST("/create",business.InsertComment)
		comments.POST("/like",business.LikeComment)
	}

	feedBack := router.Group("/help")
	{
		feedBack.POST("/inform") //举报
	}

	resource := router.Group("/resource")
	{
		//resource.StaticFS("/upload", http.Dir("/dist/images"))
		imagePath := config.GetViperString("imageSavePath")
		resource.StaticFS("/image/list", http.Dir(imagePath))
		appPath := config.GetViperString("appSavePath")
		resource.StaticFS("/apks", http.Dir(appPath))
		resource.POST("/image/dynamic/",business.GetUploadDynamicImage)
		resource.POST("/app/upload",business.UpdateApp)
		resource.GET("/app/newapp",business.GetNewApp)
	}

	area := router.Group("/area")
	{
		area.GET("/list",business.GetAreas)
	}

	httpPort := config.GetViperString("httpPort")
	fmt.Println("httpPort:", httpPort)
	router.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":"+httpPort)
}
