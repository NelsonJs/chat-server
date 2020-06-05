package normal

import (
	"chat/db/mysql_serve"
	"fmt"
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

func UploadImg(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	files := form.File["upload"]
	for _, file := range files {
		c.SaveUploadedFile(file, "D:/GoWork/images/")
	}
	c.String(http.StatusOK, "上传成功")
}
