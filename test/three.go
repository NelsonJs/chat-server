package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
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
	router.StaticFile("/apk/android.apk", "/data/mywork/apks/android.apk")
	router.POST("/upload/apk",upload)
	router.GET("/apks",apks)
	router.GET("/apk/query",get)
	router.GET("/status", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{"msg":"apk存储服务器已在线！"})
	})
	router.Run(":5885")
}

type Version struct {
	Id int64 `json:"-"`
	Url string `json:"-"`
	Num int32 `json:"num"`
	Description string `json:"description"`
	Channel string `json:"channel"`
	Createtime int64 `json:"createtime"`
}

func get(ctx *gin.Context) {
	version := ctx.DefaultQuery("version","")
	if version == "" {
		ctx.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"没有新版本",
		})
	} else {
		v,err := strconv.Atoi(version)
		if err != nil {
			ctx.JSON(http.StatusOK,gin.H{
				"code":-1,
				"msg":err.Error(),
			})
		} else {
			row := db.QueryRow("select * from versions where num = ?  order by desc",v)
			v := Version{}
			err = row.Scan(&v.Id,&v.Url,&v.Description,&v.Channel,&v.Num,&v.Createtime)
			if err != nil && err != gorm.ErrRecordNotFound{
				ctx.JSON(http.StatusOK,gin.H{"code":-1,"msg":err.Error()})
				return
			} else {
				ctx.JSON(http.StatusOK,gin.H{"code":-1,"msg":"没有新版本"})
			}
			ctx.JSON(http.StatusOK,gin.H{
				"code":1,
				"data":v,
			})
		}
	}
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
		rows,err = db.Query("select * from versions where createtime < ? limit 10 order by desc",t)
		if err != nil {
			ctx.JSON(http.StatusOK,gin.H{"code":-1,"msg":err.Error()})
			return
		}
	} else {
		rows,err = db.Query("select * from versions limit 10 order by desc")
		if err != nil {
			ctx.JSON(http.StatusOK,gin.H{"code":-1,"msg":err.Error()})
			return
		}
	}
	defer rows.Close()
	data := make([]*Version,0)
	for rows.Next() {
		v := Version{}
		err := rows.Scan(&v.Id,&v.Url,&v.Description,&v.Channel,&v.Num,&v.Createtime)
		if err != nil {
			break
		}
		data = append(data,&v)
	}
	ctx.JSON(http.StatusOK,gin.H{"code":1,"data":data})
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
	fmt.Println(header.Filename,desc,channel,num)
	ary := strings.Split(header.Filename,".")
	if len(ary) < 2 {
		ctx.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"无效的app命名",
		})
		return
	}
	if len(ary[0]) == 0 {
		ctx.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"无效的app命名",
		})
		return
	}
	var sb strings.Builder
	sb.WriteString("android.")
	sb.WriteString(ary[1])
	err = ctx.SaveUploadedFile(header,"/data/mywork/apks/"+sb.String())
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
	var url string
	if channel == "android"{
		url = "http://www.9394,cool:5885/apk/android.apk"
	} else if channel == "ios" {

	}
	result,err := db.Exec("insert into versions(url,description,channel,num,createtime)values(?,?,?,?,?)",url,desc,channel,n,time.Now().Unix())
	affected,_ := result.RowsAffected()
	if affected > 0{
		ctx.JSON(http.StatusOK,gin.H{"msg":"上传成功！"})
	} else {
		ctx.JSON(http.StatusOK,gin.H{"msg":"上传失败！"})
	}
}




