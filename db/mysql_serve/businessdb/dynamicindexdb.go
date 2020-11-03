package businessdb

import (
	"chat/config"
	"chat/db/mysql_serve"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"time"
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
	Id int64 `json:"-"`
	Did string `json:"dId"`
	Cid string `json:"cid"`
	Fid string `json:"fid"`
	Pid string `json:"pid"`
	Content string `json:"content"`
	Uid string `json:"uid"`
	Nickname string `json:"nickname"`
	Replyuid string `json:"replyuid"`
	Replyname string `json:"replyname"`
	Likenum int64 `json:"likenum"`
	Liked bool `json:"liked"`
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

//获取评论的地方需要放入缓存，否则容易爆炸
func GetComments(did string) (error,[]*CommentList){
	var comments []Comments
	tx := mysql_serve.Db.Order("createtime desc").Where("did = ? and (fid is null or fid = '')",did).Find(&comments)
	list := make([]*CommentList,0)
	fmt.Println("第一级评论数量：",len(comments))
	if tx.Error == nil && comments != nil{
		fmt.Println("执行！！")
		for _,v := range comments {
			var like Likes
			tx = mysql_serve.Db.Where("cid = ?",v.Cid).First(&like)
			if tx.Error == nil || tx.Error == gorm.ErrRecordNotFound{
				if like.Liked == 1 {
					v.Liked = true
				} else {
					v.Liked = false
				}
			}
			var commentList CommentList
			var replyComments []Comments
			tx = mysql_serve.Db.Order("createtime asc").Where("did = ? and fid = ?",did,v.Cid).Find(&replyComments)
			commentList.Comment = v
			if  len(replyComments) > 0{
				fmt.Println("replyComments数量：",replyComments)
				for i := 0; i < len(replyComments); i++ {
					var like Likes
					tx = mysql_serve.Db.Where("cid = ?",replyComments[i].Cid).First(&like)
					if tx.Error == nil || tx.Error == gorm.ErrRecordNotFound{
						if like.Liked == 1 {
							replyComments[i].Liked = true
						} else {
							replyComments[i].Liked = false
						}
					}
				}
				commentList.Comments = replyComments
			}
			list = append(list, &commentList)
		}
	}
	return tx.Error,list
}

type Likes struct {
	Id int64 `json:"-"`
	Did string `json:"did"`
	Cid string `json:"cid"`
	Liked int `json:"liked"`
	Createtime int64 `json:"createtime"`
}

func LikeComment(cid string) (error, *Comments) {
	var like Likes
	tx := mysql_serve.Db.Where("cid = ?",cid).First(&like)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound{
		return tx.Error,nil
	}
	var comment Comments
	tx = mysql_serve.Db.Where("cid = ?",cid).First(&comment)
	if tx.Error != nil {
		return tx.Error,&comment
	}
	var likeNum int64
	if like.Liked == 1 {
		likeNum = comment.Likenum - 1
		comment.Liked = false
	} else {
		likeNum = comment.Likenum + 1
		comment.Liked = true
	}
	comment.Likenum = likeNum

	mysql_serve.Db.Model(&Comments{}).Where("cid = ?",cid).Update("likenum",likeNum)
	if like.Cid == "" {
		like.Cid = cid
		like.Liked = 1
		like.Createtime = time.Now().Unix()
		mysql_serve.Db.Create(&like)
	}
	return nil,&comment
}