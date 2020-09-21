package socketservice

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Uid string
	Response chan *ResponseMsg
	Send chan *Msg
	LoginTime int64
	LogoutTIme int64
	Addr string //客户端地址
}

type ResponseMsg struct {
	code int
	content string
}

func GenerateMsg(code int,content string) *ResponseMsg {
	return &ResponseMsg{code: code,content: content}
}

func NewClient(conn *websocket.Conn)  {
	client := &Client{
		Conn: conn,
		Addr: conn.RemoteAddr().String(),
	}
	go client.Read()
	go client.Write()
}

func (client *Client) Read() {
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
	for {
		select {
		case msg := <-client.Send:
			if msg.msgType == text {
				client.Conn.WriteJSON(msg)
			}
		}
	}
}




