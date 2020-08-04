package websocket

import (
	"fmt"
	"sync"
)

type ClientManager struct {
	Clients       map[*Client]bool //所有连接的用户的集合，因为登陆不代表连接了
	ClientsLock   sync.RWMutex
	Users         map[string]*Client //登陆的用户集合
	UserLock      sync.RWMutex
	Connection    chan *Client //建立连接
	Disconnection chan *Client //断开连接
	Login         chan *Client //登陆
	Broadcast     chan []byte  //广播 向所有人发送数据
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:       make(map[*Client]bool),
		Users:         make(map[string]*Client),
		Connection:    make(chan *Client, 1000),
		Disconnection: make(chan *Client, 1000),
		Login:         make(chan *Client, 1000),
		Broadcast:     make(chan []byte, 1000),
	}
	return clientManager
}

//AddUser: 登陆后添加用户到map
func (manager *ClientManager) AddUser(client *Client) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()
	manager.Users[client.UserId] = client
	for i := range manager.Users {
		client.sendTips(1, "登录成功"+i+client.UserId+"---")
	}

}

//RemonveUser: 退出登陆后，从map中移除
func (manager *ClientManager) RemonveUser(key string) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()
	delete(manager.Users, key)
	fmt.Println("退出登陆")
}

func GetUserClient(userId string) *Client {
	v, ok := clientManager.Users[userId]
	if ok {
		return v
	}
	return nil
}

//AddClients:连接成功后，加入Clients
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	manager.Clients[client] = true
	fmt.Println("建立连接咯")
}

//AddClients:连接失败后，从Clients删除
func (manager *ClientManager) removeClient(client *Client) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()
	delete(manager.Clients, client)
	fmt.Println("断开连接咯")
}

func (manager *ClientManager) Start() {
	for {
		select {
		case login := <-manager.Login: //登陆成功
			manager.AddUser(login)
		case connection := <-manager.Connection: //建立连接
			manager.AddClients(connection)
		case disConnection := <-manager.Disconnection: //断开连接
			manager.removeClient(disConnection)
		}
	}
}
