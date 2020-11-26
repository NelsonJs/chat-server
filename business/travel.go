package business

import (
	"chat/db/mysql_serve/businessdb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTravel(c *gin.Context) {
	err,data := businessdb.GetTravelList()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":1,
		"data":data,
	})
}

func PublishTravel(c *gin.Context) {
	var travel businessdb.Travel
	if err := c.ShouldBindJSON(&travel);err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	err := businessdb.PublishTravel(&travel)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"msg":"创建成功",
		})
	}
}
