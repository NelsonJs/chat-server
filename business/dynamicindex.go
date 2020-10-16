package business

import (
	"chat/db/mysql_serve/businessdb"
	"github.com/gin-gonic/gin"
	"net/http"
)



func NearDynamicList(c *gin.Context) {
	data,err := businessdb.GetDynamics()
	if err == nil {
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"data":data,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":-1,
		"msg":err.Error(),
	})
}

func GetComments(c *gin.Context) {
	did := c.Query("did")
	if did == "" {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"argument is not enough!",
		})
		return
	}
	err,data := businessdb.GetComments(did)
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

func InsertComment(c *gin.Context) {
	var comment businessdb.Comments
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	err := businessdb.InsertComments(&comment)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"msg":"create successful",
		})
	}
}
