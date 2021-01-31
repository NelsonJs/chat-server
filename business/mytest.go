package business

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test1(c *gin.Context) {
	name,b := c.GetPostForm("name")
	pwd := c.PostForm("pwd")
	fmt.Println(b,name,pwd)
}

func TestFile(ctx *gin.Context) {
	header,err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("获取文件失败：",err)
		return
	}
	err = ctx.SaveUploadedFile(header,"d://work/html/"+header.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("提交文件成功")
}
