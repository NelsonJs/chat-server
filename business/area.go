package business

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Area struct {
	Name string `json:"n"`
	Level int64 `json:"i"`
	PLevel int64 `json:"p"`
	FirstWord string `json:"y"`
}

func GetAreas(c *gin.Context) {
	byt,err := ioutil.ReadFile("./area.json")
	byt = bytes.TrimPrefix(byt,[]byte("\xef\xbb\xbf"))
	var data []*Area
	err = json.Unmarshal(byt,&data)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-2,
			"msg":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":1,
		"data":data,
	})
}
