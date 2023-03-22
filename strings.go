package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type RandBy string

const (
	Number RandBy = "0123456789"
	Lower  RandBy = "abcdefghijklmnopqrstuvwxyz"
	Upper  RandBy = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandStr 生成随机字符串
func RandStr(length int, by ...RandBy) string {
	var chars string
	if len(by) == 0 {
		chars = fmt.Sprint(Number, Lower, Upper)
	} else {
		for _, item := range by {
			chars += string(item)
		}
	}
	s := rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	var b bytes.Buffer
	for i := 0; i < length; i++ {
		b.WriteByte(chars[s.Int63()%int64(len(chars))])
	}
	return b.String()
}
