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
	Login chan *Client
	Register chan *Client
	Broadcast chan []byte
	sync.RWMutex
}
var socketManager *manager
func init() {
	socketManager = &manager{
		Clients: make(map[string]*Client),
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
			v.LoginTime = time.Now().Unix()
			v.Response <- GenerateMsg(constants.OK,"登录成功!")
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
	m.RWMutex.Lock()
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
	m.RWMutex.Unlock()
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
