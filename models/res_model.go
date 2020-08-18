package models

import (
	"encoding/json"
)

type Res struct {
	Code int64
	Msg  string
}

type ResMsg struct {
	SendId    string
	ReceiveId string
	MsgType   int
	Content   string
}

type ResDynamic struct {
	Id               int64            `json:"id"`
	Uid              int64            `json:"uid"`
	Title            string           `json:"title"`
	Description      string           `json:"desc"`
	Img              string           `json:"img"`
	Gender           int              `json:"gender"`
	Begin            string           `json:"begin"`
	Loc              string           `json:"location"`
	Lng              float32          `json:"lng"`
	Lat              float32          `json:"lat"`
	People_num       int8             `json:"peoplenum"`
	People_total_num int8             `json:"peopletotalnum"`
	Like             int              `json:"like"`
	Comment_num      int              `json:"commentnum"`
	Comment_id       int64            `json:"commentid"`
	Time             int64            `json:"time"`
	Res_img          *json.RawMessage `json:"res_img"`
	Liked            bool             `json:"liked"`
	UserInfo_
}

type LoveIntro struct {
	Id int64 `json:"id"`
	Uid int64 `json:"uid"`
	Nickname string `json:"nickname"`
	Img string `json:"img"`
	Gender int `json:"gender"`
	Habit string `json:"habit"`
	JiGuan string `json:"ji_guan"`
	CurLocal string `json:"cur_local"`
	XueLi string `json:"xue_li"`
	Job string `json:"job"`
	ShenGao string `json:"shen_gao"`
	TiZhong string `json:"ti_zhong"`
	LoveWord string `json:"love_word"`
	CreateTime int64 `json:"create_time"`
}
