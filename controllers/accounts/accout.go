package accounts

import (
	"chat/db/mysql_serve"
	"chat/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
)

type Registers struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

func Register(c *gin.Context) {
	var register Registers
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	code, msg := mysql_serve.Register(register.Username, register.Pwd)
	if code == -1 {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"uid":      code,
			"username": register.Username,
			"pwd":      register.Pwd,
		})
	}

}

func Login(c *gin.Context) {
	var register Registers
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	code, msg := mysql_serve.Login(register.Username, register.Pwd)
	if code == -1 || code == -2 {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"uid":      code,
			"username": register.Username,
			"pwd":      register.Pwd,
		})
	}
}

func ModifyInfo(c *gin.Context) {
	var user models.UserInfo_
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"解析异常",
		})
		return
	}
	if user.Uid == "" {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"uid不存在",
		})
		return
	}
	code,err := mysql_serve.UpdateUser(user.Uid,user.Nick_name,user.Phone,user.Gender)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
	})
}

func UploadAvatar(c *gin.Context) {
	var uid = c.Query("uid")
	fmt.Printf("uid--->%s\n",uid)
	head,err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	fmt.Printf("头像名称：%s \n",head.Filename)
	filename := filepath.Base(head.Filename)
	var index = strings.LastIndex(filename,".")
	var mName = uid+filename[index:]
	var path string
	if err = c.SaveUploadedFile(head,"D:/GoWork/images/" + mName); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err:%s", err.Error()))
		return
	} else {
		path = "http://192.168.1.6:8080/resource/upload/"+mName
	}
	code,err := mysql_serve.UploadAvatar(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
	})
}
