package websocket

import (
	"chat/models"
	"encoding/json"
	"fmt"
	"sync"
)

type Handler func(client *Client, reqModel *models.Req)

var (
	handlers      = make(map[string]Handler)
	handerRWMutex sync.RWMutex
)

//Register: typeStr对应cmd类型
func Register(typeStr string, handler Handler) {
	handerRWMutex.Lock()
	defer handerRWMutex.Unlock()
	if typeStr == "" {
		return
	}
	handlers[typeStr] = handler
}

func getHandler(cmd string) (Handler, bool) {
	handerRWMutex.Lock()
	defer handerRWMutex.Unlock()
	v, ok := handlers[cmd]
	return v, ok
}

func processMsg(client *Client, msg []byte) {
	defer func() {
		if r := recover(); r != nil {
			//todo 异常
		}
	}()
	reqModel := &models.Req{}
	err := json.Unmarshal(msg, reqModel)
	if err != nil {
		fmt.Println("处理数据出错：", err)
		res := &models.Res{Code: -1, Msg: "数据不合法"}
		resData, _ := json.Marshal(res)
		client.SendMsg(resData)
		return
	}
	reqModel.MsgByte = msg
	if value, ok := getHandler(reqModel.Cmd); ok {
		value(client, reqModel)
	} else {
		res := &models.Res{Code: -1, Msg: "未知的消息类型"}
		resData, _ := json.Marshal(res)
		client.SendMsg(resData)
		return
	}
}
