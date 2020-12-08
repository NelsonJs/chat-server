package businessdb

import (
	"chat/db/mysql_serve"
	"errors"
	"gorm.io/gorm"
)

type Love struct {
	Id int64 `json:"-"`
	Uid string `json:"uid"`
	Img string `json:"img"`
	Title string `json:"title"`
	Name string `json:"name"`
	Gender int `json:"gender"`
	Likenum int64 `json:"likenum"`
	Status int `json:"status"`
	Createtime int64 `json:"createtime"`
}

func GetAllLovers(gender int) (error,[]*Love) {
	var loves []*Love
	tx := mysql_serve.Db.Where("gender = ?",gender).Find(&loves)
	if tx.Error == gorm.ErrRecordNotFound {
		return errors.New("暂无数据"),nil
	}
	return tx.Error,loves
}

func PublishLove(love *Love) error {
	tx := mysql_serve.Db.Create(&love)
	return tx.Error
}
