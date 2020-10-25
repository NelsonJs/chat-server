package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
)

func Md5(content string) string {
	data := []byte(content)
	has := md5.Sum(data)
	return fmt.Sprintf("%x",has)
}

func Md5WithTime(content string) string {
	var builer strings.Builder
	builer.WriteString(content)
	builer.WriteString(time.Now().String())
	data := []byte(builer.String())
	has := md5.Sum(data)
	return fmt.Sprintf("%x",has)
}
