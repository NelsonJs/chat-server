package accounts

import (
	"chat/db/mysql_serve"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Registers struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

func Register(c *gin.Context) {
	var register Registers
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	code, msg := mysql_serve.Register(register.Username, register.Pwd)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
