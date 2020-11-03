package business

import (
	"chat/db/mysql_serve"
	"chat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var u mysql_serve.User
	if err := c.ShouldBindJSON(&u);err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": -1,
			"msg":err.Error(),
		})
		return
	}
	if u.Phone == "" {
		c.JSON(http.StatusOK,gin.H{
			"code": -1,
			"msg":"please input phone number",
		})
		return
	}
	u.Uid = utils.Md5WithTime(u.Phone)
	u.Create_time = time.Now().Unix()
	u.Login_time = time.Now().Unix()
	err := mysql_serve.RegisterToDb(&u)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": -1,
			"msg":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code": 1,
			"data":u,
		})
	}
}

func Login(c *gin.Context) {
	var u mysql_serve.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": -1,
			"msg":err.Error(),
		})
		return
	}
	if u.Phone == "" {
		c.JSON(http.StatusOK,gin.H{
			"code": -1,
			"msg":"phone number is empty",
		})
	} else {
		u.Login_time = time.Now().Unix()
		var user *mysql_serve.User
		err,user := mysql_serve.LoginToDb(&u)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code": -1,
				"msg":err.Error(),
			})
		} else {
			c.JSON(http.StatusOK,gin.H{
				"code": 1,
				"data":user,
			})
		}
	}
}
