package business

import (
	"chat/db/mysql_serve/businessdb"
	"chat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

func InsertDynamic(c *gin.Context) {
	var dy businessdb.Dynamics
	if err := c.ShouldBindJSON(&dy); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"data":err.Error(),
		})
		return
	}
	if dy.Uid == "" {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"data":"参数缺失",
		})
		return
	}
	dy.Createtime = time.Now().Unix()
	dy.Did = utils.Md5WithTime(dy.Uid)
	businessdb.InsertDynamic(&dy)
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
