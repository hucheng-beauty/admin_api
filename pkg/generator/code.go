package generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// SmsCode 生成width长度的短信验证码
func SmsCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// RandomString 随机生成指定长度 数字+大小写字母 的字符串
func RandomString(width int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := []byte(str)
	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < width; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
