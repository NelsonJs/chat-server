package models

type LoginModel struct {
	AppId  uint32
	UserId string
	Sig    string
}

type UserInfo struct {
	Uid      string `json:"uid"`
	UserName string `json:"nick_name"`
	Phone string `json:"phone"`
	Gender   string    `json:"gender"`
}

type UserInfo_ struct {
	Uid      string `json:"uid"`
	Nick_name string `json:"nick_name"`
	Phone string `json:"phone"`
	Gender   string    `json:"gender"`
}
