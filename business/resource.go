package business

import (
	"chat/config"
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
		iPath := config.GetViperString("imagePathIp")
		httpPort := config.GetViperString("httpPort")
		imageSavePath := config.GetViperString("imageSavePath")
		var url = "http://"+iPath+":"+httpPort+"/resource/image/list/"+file.Filename
		err := c.SaveUploadedFile(file,imageSavePath+file.Filename)
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
