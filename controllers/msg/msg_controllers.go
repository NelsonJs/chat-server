package msg

import (
	"fmt"

	"chat/db/mysql_serve"

	"github.com/gin-gonic/gin"
)

func SendTxtMsg(c *gin.Context) {
	uid := c.PostForm("uid")
	msg := c.PostForm("msg")
	fmt.Println(uid, "发送消息：", msg)
}

func GetConversations(c *gin.Context) {
	uid := c.Query("uid")
	mysql_serve.GetConversations(uid)
}
