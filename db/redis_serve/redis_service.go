package redis_serve

import (
	"github.com/gomodule/redigo/redis"
)

var (
	conn redis.Conn
	err  error
)

func init() {
	conn, err = redis.Dial("tcp", "tredis:6379")
	if err != nil {
		panic(err)
	}
}
