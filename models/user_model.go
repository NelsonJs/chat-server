package models

type LoginModel struct {
	AppId  uint32
	UserId string
	Sig    string
}

type UserInfo struct {
	Uid      string `json:"uid"`
	UserName string `json:"nick_name"`
	Phone    string `json:"phone"`
	Gender   int `json:"gender"`
}

type UserInfo_ struct {
	Uid       string `json:"uid"`
	Nick_name string `json:"nickname"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	Pwd string `json:"pwd"`
	Avatar string `json:"avatar"`
	Create_time int64 `json:"createTime"`
	Login_time int64 `json:"loginTime"`
	Logout_time int64 `json:"logout_time"`
	Status int `json:"status"`
	Year_old int `json:"year_old"`
}
