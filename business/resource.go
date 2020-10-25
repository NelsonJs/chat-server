package business

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUploadDynamicImage(c *gin.Context) {
	mf,err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	files := mf.File["uploads"]
	urls := make([]string,0)
	for _,file := range files {
		var url = "http://192.168.0.109:8080/resource/image/list/"+file.Filename
		err := c.SaveUploadedFile(file,"D:/GoWork/active_img/"+file.Filename)
		if err != nil{
			c.JSON(http.StatusOK,gin.H{
				"code":-1,
				"msg":err.Error(),
			})
			return
		}
		urls = append(urls, url)
	}
	c.JSON(http.StatusOK,gin.H{
		"code":1,
		"data":urls,
	})
}
