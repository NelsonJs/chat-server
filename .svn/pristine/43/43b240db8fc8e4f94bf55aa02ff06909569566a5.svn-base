package routers

import (
	"chat/service/websocket"
)

func InitScocketRouters() {
	websocket.Register("register", websocket.RegisterController)
	websocket.Register("login", websocket.LoginController)
	websocket.Register("sendTxt", websocket.SendTxt)
}
