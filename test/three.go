package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/pw", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{"msg":"老婆我爱你啊~~~"})
	})
	router.Run(":8080")
}
