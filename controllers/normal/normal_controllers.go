package normal

import (
	"chat/db/mysql_serve"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NearDynamic(c *gin.Context) {
	err, list := mysql_serve.NearDynamic()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": list,
	})
}
