package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

type Msg2 struct {
	Cmd string //动作类型 login logout register Send
	Uid string
	PeerUid string
	Nickname string
	Avatar string
	Gender int
	Content string //文本内容
	MsgType int //自定义的消息类型，跟websocket的不同
	AppId int //1-android 2-ios 3-web
}

func main() {
	conn,_,err := websocket.DefaultDialer.Dial("ws://127.0.0.1:6767/serveWs",nil)
	if err != nil {
		fmt.Println("连接ws错误",err.Error())
		return
	}
	msg := Msg2{
		Cmd: "send",
		Uid: "101",
		PeerUid: "100",
		Nickname: "huli",
		Avatar: "",
		Gender: 2,
		Content: "这是two发送来的内容",
		MsgType: 1,
		AppId: 1,
	}
	go func() {
		for {
			_,byt,err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("客户端B接收到消息：%s\n",byt)
		}
	}()
	msg.Cmd = "register"
	byt,_ := json.Marshal(msg)
	err = conn.WriteMessage(websocket.TextMessage,byt)
	if err != nil {
		fmt.Println(err)
	}
	for {
		fmt.Println("客户端B：请输入内容..")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		msg.Cmd = "send"
		msg.Content = sc.Text()
		byt,_ := json.Marshal(&msg)
		err = conn.WriteMessage(websocket.TextMessage,byt)
		if err != nil {
			fmt.Println(err)
		}
	}
}
