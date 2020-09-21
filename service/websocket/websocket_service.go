package websocket

import (
	"chat/db/redis_serve"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

var (
	clientManager = NewClientManager()
	appIds        = []uint32{101, 102}
	redisManager  *redis_serve.RedisManager
)

func StartWebSocket(redisM *redis_serve.RedisManager) {
	redisManager = redisM
	socketPort := viper.GetString("webSocketPort")
	http.HandleFunc("/serveWs", serveWs)
	go clientManager.Start()
	http.ListenAndServe(":"+socketPort, nil)
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
	curTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, curTime, redisManager)
	go client.read()
}
