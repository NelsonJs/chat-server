package business

import (
	"chat/config"
	"chat/db/mysql_serve/businessdb"
	"chat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
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

func UpdateApp(c *gin.Context) {
	header,err := c.FormFile("app")
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	_version := c.DefaultQuery("version","0")
	v,err := strconv.Atoi(_version)
	if v <= 0 {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"请设置正确的版本号！",
		})
		return
	}
	desc := c.DefaultQuery("description","")
	channel := c.DefaultQuery("channel","")
	if channel != "android" || channel != "ios" {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"仅支持安卓跟苹果系统！",
		})
		return
	}
	appSavePath := config.GetViperString("appSavePath")
	var sb strings.Builder
	sb.WriteString("hometwon_app")
	sb.WriteString(time.Now().String())
	sb.WriteString(header.Filename)
	name := appSavePath+sb.String()
	err = c.SaveUploadedFile(header,name)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	err = businessdb.UpdateApk(channel,name,desc,v)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"msg":"发布新版本成功",
		})
	}
}

func GetNewApp(c *gin.Context) {
	var apk businessdb.Apk
	if err := c.ShouldBindJSON(&apk); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	if apk.Channel != "android" || apk.Channel != "ios" {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"仅支持安卓跟苹果系统！",
		})
		return
	}
	if apk.Num <= 0 {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"请设置正确的版本号！",
		})
		return
	}
	err,apkData := businessdb.GetApk(apk.Channel,apk.Num)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"data":apkData,
		})
	}
}
