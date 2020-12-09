package businessdb

import (
	"chat/config"
	"chat/db/mysql_serve"
)

type Apk struct {
	Name string `json:"name"`
	Num int `json:"version"`
	Description string `json:"description"`
	Channel string `json:"channel"` //android ios
	Createtime int64 `json:"createtime"`
}

func UpdateApk(channel,name,description string,version int) error {
	var apk = Apk{Channel: channel,Name: name,Description:description,Num: version}
	tx := mysql_serve.Db.Create(&apk)
	return tx.Error
}

func GetApk(channel string,version int) (error,*Apk) {
	var apk Apk
	tx := mysql_serve.Db.Order("createtime desc").Where("channel = ? and num > ?",channel,version).First(&apk)
	if tx.Error != nil {
		return tx.Error,&apk
	} else {
		iPath := config.GetViperString("imagePathIp")
		httpPort := config.GetViperString("httpPort")
		url := "http://"+iPath+":"+httpPort+"/resource/app/"+apk.Name
		apk.Name = url
		return nil,&apk
	}
}