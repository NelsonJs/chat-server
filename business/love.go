package business

import (
	"chat/db/mysql_serve/businessdb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetLoveAll(c *gin.Context) {
	gender := c.DefaultQuery("gender","0")
	g,err := strconv.Atoi(gender)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	err,data := businessdb.GetAllLovers(g)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"data":data,
		})
	}
}

func PublishLove(c *gin.Context) {
	var love businessdb.Love
	if err := c.ShouldBindJSON(&love); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	err := businessdb.PublishLove(&love)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"msg":"发布成功",
		})
	}
}
