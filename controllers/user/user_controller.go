package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SendTxtMsg(c *gin.Context) {
	uid := c.PostForm("uid")
	msg := c.PostForm("msg")
	fmt.Println(uid, "发送消息：", msg)
}
