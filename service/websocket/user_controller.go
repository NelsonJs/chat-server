package websocket

// func Register("login")

import (
	"chat/models"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//RegisterController: 注册
func RegisterController(client *Client, reqModel *models.Req) {
	if client == nil || reqModel == nil {
		return
	}
	//先根据手机号查看是否被注册
	//可用则加入输入库
	//reqModel.UserId = reqModel.PhoneNumber //暂时代替userId
	v, err := client.Redis.Redis.Do("EXISTS", reqModel.UserId)
	if err != nil {
		return
	}
	is_key_exit, _ := redis.Bool(v, err)
	fmt.Println("000是否存在：", is_key_exit)
	if !is_key_exit {
		vmap := map[string]interface{}{"appId": reqModel.AppId, "userId": reqModel.UserId, "status": 1}
		client.Redis.Redis.Do("SET", v, vmap)
		client.sendTips(1, "注册成功")
		ss, _ := redis.String(client.Redis.Redis.Do("GET", v))
		fmt.Println("注册成功：", ss)
		return
	} else {
		client.sendTips(-1, "该用户已经存在")
		return
	}
	return
}

//LoginController: 暂时直接保存
func LoginController(client *Client, reqModel *models.Req) {
	client.UserId = reqModel.UserId
	fmt.Println("userId-->" + client.UserId)
	clientManager.Login <- client
}

func GetUserInfo(client *Client, reqModel *models.Req) {

}

//UserIsRegister: 检测用户是否注册
func UserIsRegister(client *Client, reqModel *models.Req) bool {
	// v, err := client.Redis.Redis.Do("EXISTS", reqModel.UserId)
	// fmt.Println("是否存在：", v)
	// if err != nil {
	// 	fmt.Println("查询redis key是否存在出错：", err)
	// }
	// ss, _ := redis.String(v, err)
	// is_key_exists, _ := redis.Bool(v, err)
	// fmt.Println("是否存在--：", is_key_exists, "值-》", ss)
	// return is_key_exists

	return true
}
