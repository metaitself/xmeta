package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func RandomInt(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Intn(max-min)
}

func RandomIntN(length int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成随机数 长度为length
	// 比如length为6，那么生成的随机数为100000~999999之间的数
	// start = 10 e length-1次方
	start := math.Pow10(length - 1)
	// end = 10 e length次方 - 1
	end := math.Pow10(length) - 1
	return r.Intn(int(end)-int(start)) + int(start)
}

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)

	var ret []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		ret = append(ret, bytes[r.Intn(len(bytes))])
	}

	return string(ret)
}

func GetRandomVerifyCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", r.Int31n(1000000))
}
