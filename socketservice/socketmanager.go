package socketservice

import (
	"chat/config"
	"chat/constants"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type manager struct {
	Clients map[string]*Client //存储所有注册过的用户
	Online map[string]*Client //在线用户
	Login chan *Client
	Register chan *Client
	Broadcast chan []byte
	sync.RWMutex
}
var socketManager *manager
func init() {
	socketManager = &manager{
		Clients: make(map[string]*Client),
		Online: make(map[string]*Client),
		Login: make(chan *Client,10),
		Register: make(chan *Client,10),
		Broadcast: make(chan []byte,10),
	}
}
//监听事件
func listenChanEvent() {
	for {
		select {
			case loginClient := <- socketManager.Login:
				socketManager.userlogin(loginClient)
			case registerClient := <- socketManager.Register:
				socketManager.userRegister(registerClient)
		}
	}
}

//登录
func (m *manager) userlogin(client *Client) {
	m.RWMutex.Lock()
	if client != nil && client.Uid != ""{
		v,ok := m.Clients[client.Uid]
		if ok {
			//判断是否已经登录过，已登录则跳过
			vLogin,okLogin := m.Online[client.Uid]
			if okLogin { //已经是登录状态
				vLogin.Response <- GenerateMsg(constants.OK,"已经是登录状态!")
			} else {
				m.Online[client.Uid] = v //设为登录
				go m.heartBeat(client)//开启心跳维持
				v.LoginTime = time.Now()
				v.Response <- GenerateMsg(constants.OK,"登录成功!")
			}
		} else {
			client.Response <- GenerateMsg(constants.ErrNotRegister,"该用户暂未注册!")
		}
	} else {
		client.Response <- GenerateMsg(constants.ErrParameters,"登录失败，缺少uid!")
	}
	m.RWMutex.Unlock()
}


//注册
func (m *manager) userRegister(client *Client) {
	if client != nil && client.Uid != ""{
		fmt.Println("开始注册..")
		_,ok := socketManager.Clients[client.Uid]
		if ok {
			//已经注册过
			client.Response <- GenerateMsg(constants.ErrHasRegistered,"该用户已被注册!")
		} else {
			socketManager.Clients[client.Uid] = client
			client.Response <- GenerateMsg(constants.OK,"注册成功!")
		}
	} else {
		client.Response <- GenerateMsg(constants.ErrParameters,"注册失败，缺少uid!")
	}
}

func (m *manager) heartBeat(client *Client) {
	client.Conn.SetReadDeadline(time.Now().Add(client.PoneTime*time.Second))
	client.Conn.SetPongHandler(func(appData string) error {
		client.Conn.SetReadDeadline(time.Now().Add(client.PoneTime*time.Second))
		if appData != "" {
			 fmt.Println("handler-->",appData)
		} else {
			fmt.Printf("搞不好%s掉线了\n",client.Uid)
		}
		return nil
	})
	t := time.NewTicker(4*time.Second)
	count := 0
	for {
		select {
		case <-t.C:
			fmt.Println("正在ping")
			err := client.Conn.WriteMessage(websocket.PingMessage,[]byte("服务端发起ping连接"))
			if err != nil {
				count ++
				if count >= 3 {
					fmt.Println("终止此循环")
					return
				}
				fmt.Println("ping发送失败：",err)
			}
		}
	}
}

func StartSocket() {
	go listenChanEvent()
	port := config.GetViperString("webSocketPort")
	http.HandleFunc("/serveWs", serveWs)
	err := http.ListenAndServe(":"+port, nil)
	fmt.Println(err)
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}).Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("连接错误：", err.Error())
		http.NotFound(w, r)
		return
	}
	fmt.Println("连接成功")
	NewClient(conn)
}
