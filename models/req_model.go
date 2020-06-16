package models

type Req struct {
	Cmd string
	Uid string
	Message
}

type PublishDynamic struct {
	Uid string `json:"uid"`
	Title string `json:"title"`
	Ids   []int64 `json:"ids"`
}
