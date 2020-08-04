package websocket

import (
	"chat/db/mysql_serve"
)


//IsRegister： 判断是否注册
func IsRegister(client *Client) bool {
	//1.Redis中查找
	_, err := client.Redis.Redis.Do("EXISTS", client.UserId)
	if err != nil {
		//2.mysql中查找
		uid := mysql_serve.GetChatUid(client.UserId)
		if uid != "" {
			return true
		}
		return false
	}
	return true
}

func IsRegisterWithUid(client *Client, uid string) bool {
	//1.Redis中查找
	_, err := client.Redis.Redis.Do("EXISTS", uid)
	if err != nil {
		//2.mysql中查找
		uid := mysql_serve.GetChatUid(uid)
		if uid != "" {
			return true
		}
		return false
	}
	return true
}
