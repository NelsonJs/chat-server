package routers

import (
	"chat/controllers/accounts"
	"chat/controllers/msg"
	"chat/controllers/normal"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Init() {
	router := gin.Default()
	user_ := router.Group("/user")
	{
		user_.GET("/conversations", msg.GetConversations)
		user_.POST("/register", accounts.Register)
		user_.POST("/login", accounts.Login)
		user_.GET("/record", msg.GetChatRecord)
	}
	index := router.Group("/index")
	{
		index.GET("/neardynamic", normal.NearDynamic)
	}
	// msg := router.Group("/msg")
	// {
	// 	//msg.POST("/sendTxtMsg")
	// }
	httpPort := viper.GetString("app.httpPort")
	fmt.Println("httpPort:", httpPort)
	router.Run(":8080")
}
