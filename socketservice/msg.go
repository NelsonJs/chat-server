package socketservice

import (
	"chat/config"
	"chat/constants"
	"chat/db/mysql_serve"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type Msg struct {
	Cmd      string //动作类型 login logout register Send
	Uid      string
	PeerUid  string
	Nickname string
	Avatar   string
	Gender   int
	Content  string //文本内容
	MsgType  int    //自定义的消息类型，跟websocket的不同
	Ctype    string
	AppId    int //1-android 2-ios 3-web
}

//解析接收到的消息
func Parse(client *Client, msgType int, data []byte) {
	if msgType == websocket.TextMessage {
		var msg Msg
		err := json.Unmarshal(data, &msg)
		client.Uid = msg.Uid
		if err == nil {
			if msg.Cmd == "login" {
				socketManager.Login <- client
			} else if msg.Cmd == "logout" {

			} else if msg.Cmd == "register" {
				socketManager.Register <- client
			} else if msg.Cmd == "send" {
				msg.sendMsg(client)
			}
		} else {
			config.Log.Error(err.Error())
			client.Response <- GenerateMsg(constants.ErrParseData, "解析客户端消息体错误")
		}

	} else if msgType == websocket.PingMessage {

	} else if msgType == websocket.PongMessage {

	} else if msgType == websocket.BinaryMessage {

	}
}

//发消息
func (msg *Msg) sendMsg(client *Client) {
	//1.将接收到的消息放入数据表（数据漫游）
	mysql_serve.InsertMsg(&mysql_serve.Msg{Uid: msg.Uid, Peerid: msg.PeerUid, Ctype: msg.Ctype, Content: msg.Content, Create_time: time.Now().Unix()})
	//判断对方是否在线，不在线则转入离线消息（在client写入的时做，那个时候通过发送失败的错误原因判断转入离线）
	fmt.Println(*msg)
	if msg == nil || msg.PeerUid == "" {
		client.Response <- GenerateMsg(constants.ErrParameters, "请指定接收方！")
	} else {
		v, ok := socketManager.Clients[msg.PeerUid]
		if ok {
			v.Send <- msg
		} else { //接收方不存在
			client.Response <- GenerateMsg(constants.ErrNotReciver, "接收方不存在！")
		}
	}
}
