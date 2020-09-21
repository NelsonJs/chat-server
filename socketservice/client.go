package socketservice

import (
	"encoding/json"
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
	Code int
	Content string
}

func GenerateMsg(code int,content string) *ResponseMsg {
	return &ResponseMsg{Code: code,Content: content}
}

func NewClient(conn *websocket.Conn)  {
	client := &Client{
		Conn: conn,
		Response: make(chan *ResponseMsg,10),
		Send: make(chan *Msg,10),
		Addr: conn.RemoteAddr().String(),
	}
	go client.Read()
	go client.Write()
}

func (client *Client) Read() {
	for {
		msgType,byt,err := client.Conn.ReadMessage()
		if err == nil {
			var msg Msg
			err := json.Unmarshal(byt,&msg)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("--------->",msg)
			}
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
			if msg.MsgType == text {
				byt,_ := json.Marshal(*msg)
				client.Conn.WriteMessage(websocket.TextMessage,byt)
			}
		case responseMsg := <-client.Response:
			fmt.Println("write responseMsg :",*responseMsg)
			byt,_ := json.Marshal(*responseMsg)
			client.Conn.WriteMessage(websocket.TextMessage,byt)
		}
	}
}




