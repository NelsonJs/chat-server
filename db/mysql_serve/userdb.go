package mysql_serve

import (
	"chat/constants"
	"gorm.io/gorm"
)

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
	pwd := user.Pwd
	phone := user.Phone
	if pwd == "" || phone == ""{
		return constants.ErrArgumentNotExists
	}
	var u User
	tx := Db.Where("phone = ?",user.Phone).First(&u)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound{
		return tx.Error
	} else if u.Uid != "" {
		return constants.ErrUserHasRegister
	}
	tx = Db.Create(user)
	return tx.Error
}


func LoginToDb(user *User) (error,*User) {
	pwd := user.Pwd
	phone := user.Phone
	if pwd == "" || phone == ""{
		return constants.ErrArgumentNotExists,nil
	}
	var u User
	tx := Db.Where("phone = ? and pwd = ?",user.Phone,user.Pwd).First(&u)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return constants.ErrUserNotExists,nil
		}
		return tx.Error,nil
	} else if u.Uid == "" {
		return constants.ErrUserNotExists,nil
	}
	tx = Db.Model(user).Where("phone = ?",user.Phone).Update("login_time",user.Login_time)
	return tx.Error,&u
}

func QueryUser(uid string) *User {
	if uid == ""{
		return nil
	}
	var u User
	tx := Db.Where("uid = ?",uid).First(&u)
	if tx.Error != nil || tx.Error == gorm.ErrRecordNotFound{
		return nil
	}
	return &u
}

func UserIsExists(uid string) bool {
	if uid == ""{
		return false
	}
	var u User
	tx := Db.Where("uid = ?",uid).First(&u)
	if tx.Error != nil {
		return false
	}
	if u.Uid != "" {
		return true
	}
	return false
}
