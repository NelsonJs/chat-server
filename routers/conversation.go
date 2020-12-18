package routers

import (
	"chat/db/mysql_serve"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//获取会话列表
func ConversationList(c *gin.Context) {
	uid := c.Query("uid")
	fmt.Println("uid-->", uid)
	if uid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "uid为空",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": mysql_serve.GetConversations(uid),
	})
}




//获取聊天记录
func ChatRecords(c *gin.Context) {
	uid := c.Query("uid")
	peerId := c.Query("peerId")
	ctype := c.Query("ctype")

	if uid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "uid不存在",
		})
		return
	}
	if peerId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "peerId不存在",
		})
		return
	}
	if ctype == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请指明会话类型",
		})
		return
	}
	data, err := mysql_serve.GetMsgList(peerId, uid, ctype, 10, time.Now().Unix())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": data,
	})
}

//删除会话
func DelConversation(c *gin.Context) {

}

//撤回消息
func RevokeMsg(c *gin.Context) {

}
