package websocket

// func Register("login")

import (
	"chat/models"
	"fmt"
)

//LoginController: 暂时直接保存
func LoginController(client *Client, reqModel *models.Req) {
	client.UserId = reqModel.Uid
	fmt.Println("userId-->" + client.UserId)
	clientManager.Login <- client
}

func GetUserInfo(client *Client, reqModel *models.Req) {

}

//UserIsRegister: 检测用户是登录
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
