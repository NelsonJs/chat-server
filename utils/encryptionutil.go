package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(content string) string {
	data := []byte(content)
	has := md5.Sum(data)
	return fmt.Sprintf("%x",has)
}
