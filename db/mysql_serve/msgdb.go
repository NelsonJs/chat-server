package mysql_serve

import (
	"chat/config"
	"gorm.io/gorm"
)

type Msg struct {
	MsgId       string `json:"msg_id"`
	Uid         string `json:"uid"`
	Peerid      string `json:"peerid"`
	Ctype       string `json:"ctype"`
	Content     string `json:"content"`
	Msg_type    int    `json:"msg_type"`
	Pic         string
	Status      int
	Create_time int64
}

func InsertMsg(msg *Msg) (bool, error) {
	tx := db.Create(msg)
	if tx.Error != nil {
		config.Log.Error(tx.Error.Error())
		return false, tx.Error
	}
	return true, nil
}

func GetMsgList(conversationId, uid, ctype string, limit int, createTime int64) (*[]Msg, error) {
	var msgs []Msg
	tx := db.Order("create_time desc").
		Where("peerid = ? and uid = ? and create_time < ? and ctype = ?", conversationId, uid, createTime, ctype).
		Or("peerid = ? and uid = ? and create_time < ? and ctype = ?", uid, conversationId, createTime, ctype).Limit(limit).Find(&msgs)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return nil, tx.Error
	}
	return &msgs, nil
}

func GetConversations(uid string) []*Msg {
	rows, err := db.Raw(`select conversationid,any_value(ctype) as conversationType,any_value(content) as content,any_value(create_time) as create_time 
from (
select peerid as conversationid,ctype,content,create_time from msg where uid = ? and peerid <> ? 
union 
select uid as conversationid,ctype,content,create_time from msg  where uid <> ? and peerid = ?
order by 
create_time desc)as tableother group by conversationid;`, uid, uid, uid, uid).Rows()
	data := make([]*Msg, 0)
	if err != nil {
		config.Log.Error(err.Error())
		return data
	}
	defer rows.Close()
	var conversationId string = ""
	var conversationType string = ""
	var content string = ""
	var create_time int64 = 0

	for rows.Next() {
		err = rows.Scan(&conversationId, &conversationType, &content, &create_time)
		if err != nil {
			break
		}
		msg := Msg{Uid: uid, Peerid: conversationId, Ctype: conversationType, Content: content, Create_time: create_time}
		data = append(data, &msg)
	}
	return data

}
