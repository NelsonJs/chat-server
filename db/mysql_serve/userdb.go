package mysql_serve

type User struct {
	Id int64 `json:"-"`
	Uid string `json:"uid"`
	Nickname string `json:"nickname"`
	Phone string `json:"phone"`
	Pwd string `json:"pwd"`
	Gender int `json:"gender"`
	Avatar string `json:"avatar"`
	Create_time int64 `json:"create_time"`
	Login_time int64 `json:"login_time"`
	Logout_time int64 `json:"logout_time"`
	Status int `json:"status"`
}

func RegisterToDb(user *User) error {
	tx := Db.Create(user)
	return tx.Error
}

func LoginToDb(user *User) error {
	tx := Db.Model(user).Where("uid = ?",user.Uid).Update("login_time",user.Login_time)
	return tx.Error
}
