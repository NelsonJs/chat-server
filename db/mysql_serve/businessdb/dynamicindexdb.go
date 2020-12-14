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

type DynamicsQuery struct {
	Id int64 `json:"-"`
	Did string `json:"did"`
	Uid string `json:"uid"`
	Nickname string `json:"nickname"`
	Title string `json:"title"`
	Avatar string `json:"avatar"`
	Gender int `json:"gender"`
	Likenum int64 `json:"likeNum"`
	Liked bool `json:"liked"`
	Location string `json:"location"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Createtime int64 `json:"createTime"`
	Resimg JSON `json:"resImg"`
	Description string `json:"desc"`
}

type Dynamics struct {
	Id int64 `json:"-"`
	Did string `json:"did"`
	Uid string `json:"uid"`
	Nickname string `json:"nickname"`
	Title string `json:"title"`
	Avatar string `json:"avatar"`
	Gender int `json:"gender"`
	Likenum int64 `json:"likeNum"`
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


func GetDynamics(uid string) ([]*map[string]interface{},error){
	var data []DynamicsQuery
	tx := mysql_serve.Db.Table("dynamics").Order("createtime desc").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	list := make([]*map[string]interface{},0)
	if data != nil && len(data) > 0 {
		fmt.Println(len(data))
		for _,v := range data {
			if uid != "" {
				var like Likes
				tx = mysql_serve.Db.Where("uid = ? and did = ?",uid,v.Did).First(&like)
				if tx != nil && tx.Error != gorm.ErrRecordNotFound {
					v.Liked = true
				}
			}
			err,comments := GetComments(v.Did)
			if err != nil{
				if err == gorm.ErrRecordNotFound {
					m := Struct2Map(v)
					list = append(list, &m)
				} else {
					fmt.Println(err.Error())
					config.Log.Error(err.Error())
				}
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
type CommentsCreate struct {
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
	Status int `json:"status"`
	Createtime int64 `json:"createTime"`
}

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


func InsertComments(c *CommentsCreate) error {
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
	Uid string `json:"uid"`
	Did string `json:"did"`
	Cid string `json:"cid"`
	Liked int `json:"liked"`
	Createtime int64 `json:"createtime"`
}

func LikeDynamic(uid, did string) (error,*DynamicsQuery) {
	var dy DynamicsQuery
	mysql_serve.Db.Transaction(func(tx *gorm.DB) error {
		var like Likes
		t := tx.Where("uid = ? and did = ?",uid,did).First(&like)
		if t.Error != nil && t.Error != gorm.ErrRecordNotFound{
			return t.Error
		}
		t = tx.Where("did = ?",did).First(&dy)
		if t.Error != nil {
			return t.Error
		}
		var likeNum int64
		if like.Liked == 1 {
			likeNum = dy.Likenum - 1
			dy.Liked = false
			t = tx.Where("uid = ? and did = ?",uid,did).Delete(&like)
			if t.Error != nil {
				return t.Error
			}
		} else {
			likeNum = dy.Likenum + 1
			dy.Liked = true

			like.Did = did
			like.Uid = uid
			like.Liked = 1
			like.Createtime = time.Now().Unix()
			t = tx.Create(&like)
			if t.Error != nil {
				return t.Error
			}
		}
		dy.Likenum = likeNum

		t = tx.Model(&Dynamics{}).Where("did = ?",did).Update("likenum",likeNum)
		if t.Error != nil {
			return t.Error
		}
		return nil
	})

	return nil,&dy
}

func LikeComment(uid,cid string) (error, *Comments) {
	var comment Comments
	mysql_serve.Db.Transaction(func(tx *gorm.DB) error {
		var like Likes
		t := tx.Where("uid = ? and cid = ?",uid,cid).First(&like)
		if t.Error != nil && t.Error != gorm.ErrRecordNotFound{
			return t.Error
		}
		t = tx.Where("cid = ?",cid).First(&comment)
		if t.Error != nil {
			return t.Error
		}
		var likeNum int64
		if like.Liked == 1 {
			likeNum = comment.Likenum - 1
			comment.Liked = false
			t = tx.Where("uid = ? and cid = ?",uid,cid).Delete(&like)
			if t.Error != nil {
				return t.Error
			}
		} else {
			likeNum = comment.Likenum + 1
			comment.Liked = true

			like.Cid = cid
			like.Uid = uid
			like.Liked = 1
			like.Createtime = time.Now().Unix()
			t = tx.Create(&like)
			if t.Error != nil {
				return t.Error
			}
		}
		comment.Likenum = likeNum

		t = tx.Model(&Comments{}).Where("cid = ?",cid).Update("likenum",likeNum)
		if t.Error != nil {
			return t.Error
		}
		return nil
	})

	return nil,&comment
}