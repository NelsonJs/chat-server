package msg

import (
	"fmt"
	"net/http"
	"strconv"

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
	ids := mysql_serve.GetConversations(uid)
	c.JSON(http.StatusOK, gin.H{
		"data": ids,
	})
}

type chat struct {
	selfId  int64
	otherId int64
}

func GetChatRecord(c *gin.Context) {
	// ch := chat{}
	// if err := c.ShouldBindJSON(&ch); err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": -1,
	// 		"msg":  err.Error(),
	// 	})
	// 	return
	// }
	selfId := c.PostForm("selfId")
	otherId := c.PostForm("otherId")
	sid, err := strconv.ParseInt(selfId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	oid, err := strconv.ParseInt(otherId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	code, msg, data := mysql_serve.GetChatRecord(sid, oid)
	if code == -1 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": data,
		})
	}
}
