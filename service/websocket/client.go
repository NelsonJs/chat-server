package websocket

import (
	"chat/db/redis_serve"
	"chat/models"
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/gorilla/websocket"
)

type Client struct {
	Addr          string          //客户端地址
	Socket        *websocket.Conn //websocket连接
	Send          chan []byte     //待发送的数据
	AppId         uint32          //登陆平台的id android/ios/web
	UserId        string          //用户id，登陆才有
	FirstTime     uint64          //首次连接时间
	HeartbeatTime uint64          //用户上次心跳时间
	LoginTime     uint64          //登陆时间
	Redis         *redis_serve.RedisManager
}

func NewClient(addr string, conn *websocket.Conn, firstTime uint64, redisM *redis_serve.RedisManager) (client *Client) {
	return &Client{
		Addr:          addr,
		Socket:        conn,
		Send:          make(chan []byte, 1000),
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
		Redis:         redisM,
	}
}

type Login struct {
	AppId  string
	UserId string
	Client *Client
}

func (l *Login) GetKey() (key string) {
	key = fmt.Sprintf("%s,%s", l.AppId, l.UserId)
	return
}

func (client *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("read stop", string(debug.Stack()), r)
		}
	}()
	defer func() {
		close(client.Send)
	}()

	for {
		_, message, err := client.Socket.ReadMessage()
		if err != nil {
			fmt.Println("读取消息错误：", client.Addr, err)
			return
		}
		fmt.Printf("读取消息：%s \n", message)
		processMsg(client, message)
		// client.Send <- message
		// req := models.Req{}
		// err = json.Unmarshal(message, &req)
		// if err != nil {
		// 	fmt.Println("解析json失败 ", err.Error())
		// 	break
		// }
		// fmt.Println("appid-->", req.AppId)
	}
}

func (client *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()
	defer func() {
		client.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				fmt.Println("发送数据错误")
				return
			}
			client.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (client *Client) SendMsg(byt []byte) {
	client.Send <- byt
}

func (client *Client) sendTips(code int64, msg string) {
	res := &models.Res{Code: code, Msg: msg}
	resData, _ := json.Marshal(res)
	client.SendMsg(resData)
}
