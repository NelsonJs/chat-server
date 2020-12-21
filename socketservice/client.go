package socketservice

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	Conn *websocket.Conn
	Uid string
	Response chan *ResponseMsg
	Send chan *Msg
	LoginTime time.Time
	LogoutTIme time.Time
	PoneTime time.Duration
	Addr string //客户端地址
}

type ResponseMsg struct {
	Code int
	Content string
	MsgType int
}

func GenerateMsg(code int,content string) *ResponseMsg {
	return &ResponseMsg{Code: code,Content: content}
}

func GenerateNoticeMsg(code,msgType int,content string) *ResponseMsg {
	return &ResponseMsg{Code: code,MsgType: msgType,Content: content}
}

func NewClient(conn *websocket.Conn)  {
	client := &Client{
		Conn: conn,
		Response: make(chan *ResponseMsg,10),
		Send: make(chan *Msg,10),
		PoneTime: 8,
		Addr: conn.RemoteAddr().String(),
	}
	go client.Read()
	go client.Write()
}

func (client *Client) Read() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("read 出现异常%s\n",err)
			close(client.Send)
			close(client.Response)
		}
	}()
	for {
		msgType,byt,err := client.Conn.ReadMessage()
		if err == nil {
			Parse(client,msgType,byt)
		} else {
			fmt.Println(err.Error())
		}
	}
}

func (client *Client) Write() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("write meet err:",err)
			err = client.Conn.Close()
			if err != nil {
				fmt.Printf("recover中关闭conn出错：%s\n",err)
			}
		}
	}()
	for {
		select {
		case msg := <-client.Send:
			if msg.MsgType == text {
				byt,_ := json.Marshal(*msg)
				fmt.Println(msg)
				client.Conn.WriteMessage(websocket.TextMessage,byt)
			}
		case responseMsg := <-client.Response:
			fmt.Println("正在接收消息")
			byt,_ := json.Marshal(*responseMsg)
			fmt.Println(responseMsg)
			client.Conn.WriteMessage(websocket.TextMessage,byt)
		}
	}
}




