package utils

import (
	secureRand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
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

func GetRandomUniqueDigit(length int) string {
	s := ""
	m := make(map[int]bool)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for length > 0 {
		digit := r.Intn(9)
		if _, ok := m[digit]; !ok {
			s += strconv.Itoa(digit)
			m[digit] = true
			length--
		}
	}
	return s
}

// 获取伪随机数种子
func getRand() *rand.Rand {
	return rand.New(rand.NewSource(GetSecureRandInt(time.Now().UnixNano())))
}

// GetSecureRandInt 返回一个真随机数生成[0, max)的整数
func GetSecureRandInt(max int64) int64 {
	secureMax := new(big.Int).SetInt64(max)
	res, err := secureRand.Int(secureRand.Reader, secureMax)
	if err != nil {
		fmt.Printf("Can't generate random value: %v, %v", res, err)
	}
	return res.Int64()
}

// GetSecureLenRand 返回一个真随机数按进制算len长度的整数，注意len必须要大于等于2，比如3为2位
func GetSecureLenRand(len int) int64 {
	res, err := secureRand.Prime(secureRand.Reader, len)
	if err != nil {
		fmt.Printf("Can't generate random value because the len is < 2")
	}
	return res.Int64()
}

// GetSecureRandIntValue 切片生成随机数
func GetSecureRandIntValue(n int) byte {
	val := make([]byte, n)
	//rand.Reader是一个全局、共享的密码用强随机数生成器
	n, err := secureRand.Read(val)
	if err != nil {
		fmt.Print("Can't generate securerandom value")
	}
	return val[n]
}

// GetRandomInt 取值范围在[0,n)的伪随机int值， 如果n<=0会panic
func GetRandomInt(n int) int {
	return getRand().Intn(n)
}

// GetRandomInt31n 取值范围在[0,n)的伪随机int值， 如果n<=0会panic
func GetRandomInt31n(n int32) int32 {
	return getRand().Int31n(n)
}

// GetRandomInt63n 返回一个取值范围在[0,n)的伪随机int64值， 如果n<=0会panic。
func GetRandomInt63n(n int64) int64 {
	return getRand().Int63n(n)
}

// GetRandomFloat 返回一个取值范围在[0.0, 1.0)的伪随机float32值。
func GetRandomFloat() float32 {
	return getRand().Float32()
}

// GetRandomFloat64 返回一个取值范围在[0.0, 1.0)的伪随机float64值。
func GetRandomFloat64() float64 {
	return getRand().Float64()
}

// GetContextLen 返回一个有n个元素的， [0,n)范围内整数的伪随机排列的切片。
func GetContextLen(n int) []int {
	return getRand().Perm(n)
}

// GetRandBetween 返回[min, max)范围的随机数
func GetRandBetween(min, max int) int {
	return getRand().Intn(max-min) + min
}
