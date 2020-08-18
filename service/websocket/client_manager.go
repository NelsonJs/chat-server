package websocket

import (
	"fmt"
	"sync"
)

type ClientManager struct {
	sync.RWMutex
	Clients         map[string]*Client //登陆(连接)的用户集合
	Disconnection chan *Client //断开连接
	Login         chan *Client //登陆/建立连接
	Broadcast     chan []byte  //广播 向所有人发送数据
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:       make(map[string]*Client),
		Disconnection: make(chan *Client, 1000),
		Login:         make(chan *Client, 1000),
		Broadcast:     make(chan []byte, 1000),
	}
	return clientManager
}


func GetUserClient(userId string) *Client {
	v, ok := clientManager.Clients[userId]
	if ok {
		return v
	}
	return nil
}

//AddClients:连接成功后，加入Clients
func (manager *ClientManager) AddClients(client *Client) {
	manager.Lock()
	defer manager.Unlock()
	manager.Clients[client.UserId] = client
	for i := range manager.Clients {
		client.sendTips(1, "登录成功"+i+client.UserId+"---")
	}
}

//AddClients:连接失败后，从Clients删除
func (manager *ClientManager) removeClient(client *Client) {
	manager.Lock()
	defer manager.Unlock()
	delete(manager.Clients, client.UserId)
	fmt.Println("退出登陆")
}

func (manager *ClientManager) Start() {
	for {
		select {
		case login := <-manager.Login: //登陆成功
			manager.AddClients(login)
		case disConnection := <-manager.Disconnection: //断开连接/退出登录
			manager.removeClient(disConnection)
		}
	}
}
