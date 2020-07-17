package redis_serve

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedisManager struct {
	Redis redis.Conn
}

func ConnectRedis() *RedisManager {
	c, err := redis.Dial("tcp", "tredis:6379")
	if err != nil {
		fmt.Println("Connect to redis error:", err)
		return nil
	}
	return &RedisManager{Redis: c}
}

func (manager *RedisManager) addUser(info map[string]interface{}) {
	if info == nil {
		return
	}

}
