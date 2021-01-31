package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)
var db *sql.DB
var err error
func main() {
	db,err = sql.Open("mysql","root:6678510Jk.@tcp(localhost:3306)/hometown")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()
	router.POST("/upload/apk",upload)
	router.GET("/apks",apks)
	router.GET("/status", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{"msg":"apk存储服务器已在线！"})
	})
	router.Run(":5885")
}

type Version struct {
	Id int64 `json:"-"`
	Name string `json:"name"`
	Num int32 `json:"num"`
	Description string `json:"description"`
	Channel string `json:"channel"`
	Createtime int64 `json:"createtime"`
}

func apks(ctx *gin.Context) {
	time,b := ctx.GetPostForm("time")
	var rows *sql.Rows
	var err error
	if b {
		t,err := strconv.Atoi(time)
		if err != nil {
			ctx.JSON(http.StatusOK,gin.H{"code":-1,"msg":err.Error()})
			return
		}
		rows,err = db.Query("select * from versions where createtime > ? limit 10",t)
		if err != nil {
			ctx.JSON(http.StatusOK,gin.H{"code":-1,"msg":err.Error()})
			return
		}
	} else {
		rows,err = db.Query("select * from versions limit 10")
		if err != nil {
			ctx.JSON(http.StatusOK,gin.H{"code":-1,"msg":err.Error()})
			return
		}
	}
	data := make([]*Version,0)
	for rows.Next() {
		v := Version{}
		err := rows.Scan(&v.Id,&v.Name,&v.Description,&v.Channel,&v.Num,&v.Createtime)
		if err != nil {
			break
		}
		data = append(data,&v)
	}
	ctx.JSON(http.StatusOK,gin.H{"code":1,"msg":"successful"})
}

func upload(ctx *gin.Context) {
	header,err := ctx.FormFile("file")
	desc := ctx.PostForm("desc")
	channel := ctx.PostForm("channel")
	num := ctx.PostForm("num")
	if err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	err = ctx.SaveUploadedFile(header,"/data/mywork/apks/"+header.Filename)
	if err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":err.Error(),
		})
		return
	}
	n := 0
	if num != "" {
		n,_ = strconv.Atoi(num)
	}
	result,err := db.Exec("insert into versions(name,description,channel,num,createtime)values(?,?,?,?,?)",header.Filename,desc,channel,n,time.Now().Unix())
	affected,_ := result.RowsAffected()
	if affected > 0{
		ctx.JSON(http.StatusOK,gin.H{"msg":"上传成功！"})
	} else {
		ctx.JSON(http.StatusOK,gin.H{"msg":"上传失败！"})
	}
}




