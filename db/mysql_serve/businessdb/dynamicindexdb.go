package businessdb

import (
	"chat/config"
	"chat/db/mysql_serve"
	"database/sql/driver"
	"encoding/json"
	"errors"
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
	Resimg JSON `json:"resImg"`
	Description string `json:"desc"`
}

type JSON struct {
	json.RawMessage
}

// Value get value of JSON
func (j JSON) Value() (driver.Value, error) {
	if len(j.RawMessage) == 0 {
		return nil, nil
	}
	return j.MarshalJSON()
}

// Scan scan value into JSON
func (j *JSON) Scan(value interface{}) error {
	str, ok := value.(string)
	if ok {
		bytes := []byte(str)
		return json.Unmarshal(bytes, j)
	}
	if reflect.ValueOf(value).Kind() == reflect.Slice {
		return json.Unmarshal(reflect.ValueOf(value).Bytes(),j)
	}
	return errors.New(fmt.Sprint("Failed to unmarshal JSONB value (strcast):", value))
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
	Fid string `json:"fid"`
	Pid string `json:"pid"`
	Content string `json:"content"`
	Uid string `json:"uid"`
	Nickname string `json:"nickname"`
	ReplyUid string `json:"replyuid"`
	Replyname string `json:"replyname"`
	Likenum int64 `json:"likenum"`
	Status int `json:"status"`
	Createtime int64 `json:"createTime"`
}

type CommentList struct {
	Comment Comments `json:"comment"`
	Comments []Comments `json:"comments"`
}


func InsertComments(c *Comments) error {
	tx := mysql_serve.Db.Create(c)
	return tx.Error
}

func GetComments(did string) (error,[]*CommentList){
	var comments []Comments
	tx := mysql_serve.Db.Order("createtime desc").Where("did = ? and fid is null",did).Find(&comments)
	list := make([]*CommentList,0)
	fmt.Println("第一级评论数量：",len(comments))
	if tx.Error == nil && comments != nil{
		fmt.Println("执行！！")
		for _,v := range comments {
			var commentList CommentList
			var replyComments []Comments
			tx = mysql_serve.Db.Order("createtime desc").Where("did = ? and fid = ?",did,v.Cid).Find(&replyComments)
			commentList.Comment = v
			if  len(replyComments) > 0{
				fmt.Println("replyComments数量：",replyComments)
				commentList.Comments = replyComments
			}
			list = append(list, &commentList)
		}
	}
	return tx.Error,list
}