package models

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
	Id               int64   `json:"id"`
	Uid              int64   `json:"uid"`
	Title            string  `json:"title"`
	Description      string  `json:"desc"`
	Img              string  `json:"img"`
	Gender           int     `json:"gender"`
	Begin            string  `json:"begin"`
	Loc              string  `json:"location"`
	Lng              float32 `json:"lng"`
	Lat              float32 `json:"lat"`
	People_num       int8    `json:"peoplenum"`
	People_total_num int8    `json:"peopletotalnum"`
	Like             int     `json:"like"`
	Comment_num      int     `json:"commentnum"`
	Comment_id       int64   `json:"commentid"`
}
