package main

import (
	"chat/socketservice"
)

// @title 微聊 API
// @version 1.0
// @description 不定时更新
// @BasePath
// @contact.name pjsong
// @contact.url xxx 网址
// @contact.email 18320944165@163.com 邮箱
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
func main() {
	socketservice.StartSocket()
}
