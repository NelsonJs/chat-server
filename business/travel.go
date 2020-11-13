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
