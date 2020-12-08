package business

import (
	"chat/config"
	"chat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
		filename := utils.Md5WithTime(file.Filename)
		var sb strings.Builder
		sb.WriteString(filename)
		sb.WriteString(file.Filename)
		var url = "http://"+iPath+":"+httpPort+"/resource/image/list/"+sb.String()
		err := c.SaveUploadedFile(file,imageSavePath+sb.String())
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
