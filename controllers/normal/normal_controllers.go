package normal

import (
	"chat/db/mysql_serve"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"chat/models"

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
	fmt.Printf("文件个数为：%d\n", len(files))
	paths := make([]string, 0)
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		fmt.Printf("图片名称为：%s \n", filename)
		var index = strings.LastIndex(filename, ".")
		var mName = strconv.FormatInt(time.Now().Unix(), 10) + filename[index:]
		if err := c.SaveUploadedFile(file, "/dist/images/"+mName); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err:%s", err.Error()))
			return
		} else {
			paths = append(paths, "http://192.168.1.6:8080/resource/upload/"+mName)
		}
	}
	res, err := mysql_serve.AddImg(paths)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": res,
	})
}

func UploadOneImg(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	files := form.File["upload"]
	paths := make([]string, 0)
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		fmt.Printf("图片名称为：%s \n", filename)
		if err := c.SaveUploadedFile(file, "/dist/images/"+filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err:%s", err.Error()))
			return
		} else {
			paths = append(paths, "http://192.168.1.6:8080/resource/upload/"+filename)
		}
	}
	res, err := mysql_serve.AddImg(paths)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": res,
	})
}

func PublishDynamic(c *gin.Context) {
	var addDynamic models.PublishDynamic
	if err := c.ShouldBindJSON(&addDynamic); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	code, err := mysql_serve.PublishDynamic(addDynamic.Uid, addDynamic.Title, addDynamic.Ids)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(code)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "测试~~~~~~",
	})
}
