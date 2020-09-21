package redis_serve

import "fmt"

func SaveString(key, value string, t ...int64) (bool, error) {
	var _t int64 = 0
	if len(t) > 0 {
		_t = t[0]
	}
	_, err := conn.Do("SET", key, value, "EX", _t)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetString(key string) (string, error) {
	reply, err := conn.Do("GET", key)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", reply), nil
}
