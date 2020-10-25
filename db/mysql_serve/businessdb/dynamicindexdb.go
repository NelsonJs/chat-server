package businessdb

import (
	"chat/config"
	"chat/db/mysql_serve"
	"encoding/json"
	"fmt"
	"reflect"
)

type Dynamics struct {
	Id int64 `json:"-"`
	Did string `json:"did"`
	Uid string `json:"uid"`
	Nickname string `json:"nickname"`
	Title string `json:"title"`
	Avatar string `json:"avatar"`
	Gender int `json:"gender"`
	Likenum int64 `json:"likeNum"`
	Liked int `json:"liked"`
	Location string `json:"location"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Createtime int64 `json:"createTime"`
	Resimg json.RawMessage `json:"resImg"`
	Description string `json:"desc"`
}



func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := t.NumField()-1; i >= 0; i-- {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}


func GetDynamics() ([]*map[string]interface{},error){
	var data []Dynamics
	tx := mysql_serve.Db.Order("createtime desc").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	list := make([]*map[string]interface{},0)
	if data != nil && len(data) > 0 {
		fmt.Println(len(data))
		for _,v := range data {
			err,comments := GetComments(v.Did)
			if err != nil {
				fmt.Println(err.Error())
				config.Log.Error(err.Error())
			} else {
				m := Struct2Map(v)
				m["comments"] = comments
				list = append(list, &m)
			}
		}
	}
	return list,nil
}

func InsertDynamic(dy *Dynamics) error {
	tx := mysql_serve.Db.Create(dy)
	return tx.Error
}


//comments
type Comments struct {
	Id string `json:"-"`
	Did string `json:"commentId"`
	Cid string `json:"cid"`
	Content string `json:"content"`
	Uid string `json:"uid"`
	Nickname string `json:"nickname"`
	Likenum int64 `json:"likenum"`
	Status int `json:"status"`
	Reply json.RawMessage `json:"reply"`
	Createtime int64 `json:"createTime"`
}


func InsertComments(c *Comments) error {
	tx := mysql_serve.Db.Create(c)
	return tx.Error
}

func GetComments(did string) (error,[]*Comments){
	var comments []*Comments
	tx := mysql_serve.Db.Order("createtime desc").Where("did = ?",did).Find(&comments)
	return tx.Error,comments
}