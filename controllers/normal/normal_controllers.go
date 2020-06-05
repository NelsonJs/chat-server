package normal

import (
	"chat/db/mysql_serve"
	"fmt"
	"net/http"
	"path/filepath"

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
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, "D:/GoWork/images"+filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err:%s", err.Error()))
			return
		}
	}
	c.String(http.StatusOK, "上传成功")
}
