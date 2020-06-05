package msg

import (
	"net/http"
	"strconv"

	"chat/db/mysql_serve"
	"chat/service/websocket"

	"chat/models"

	"github.com/gin-gonic/gin"
)

func SendTxtMsg(c *gin.Context) {
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	}
	b, err := websocket.SendText(strconv.FormatInt(msg.SendId, 10), strconv.FormatInt(msg.ReceiveId, 10), msg.Content)
	if b {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "发送成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	}
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

	selfId := c.Query("selfId")
	otherId := c.Query("otherId")
	sid, err := strconv.ParseInt(selfId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -11,
			"msg":  err.Error(),
		})
		return
	}
	oid, err := strconv.ParseInt(otherId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -21,
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
