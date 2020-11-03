package business

import (
	"chat/db/mysql_serve/businessdb"
	"chat/utils"
	"fmt"
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
			"msg":err.Error(),
		})
		return
	}
	if dy.Uid == "" {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"参数缺失",
		})
		return
	}
	fmt.Println(dy.Resimg)
	dy.Createtime = time.Now().Unix()
	dy.Did = utils.Md5WithTime(dy.Uid)
	err := businessdb.InsertDynamic(&dy)
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
	cid := utils.Md5WithTime(comment.Uid+comment.Did)
	comment.Cid = cid
	comment.Createtime = time.Now().Unix()
	err := businessdb.InsertComments(&comment)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"data":comment,
		})
	}
}

func LikeComment(c *gin.Context) {
	var comment businessdb.Comments
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	if comment.Cid == "" {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"cid为空",
		})
		return
	}
	err,co := businessdb.LikeComment(comment.Cid)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"data":co,
		})
	}
}
