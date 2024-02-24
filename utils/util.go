package utils

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"
	"unsafe"
)

// If 三目运算
func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

// Recover handles panic and logs stack info
func Recover() {
	if err := recover(); err != nil {
		const size = 64 << 10
		buf := make([]byte, size)
		buf = buf[:runtime.Stack(buf, false)]
		fmt.Printf("runtime error: %v\ntraceback:\n%v\n", err, *(*string)(unsafe.Pointer(&buf)))
	}
}

// Safe wraps a function-calling with panic recovery
func Safe(call func()) {
	defer Recover()
	call()
}

func Try(callback func() error) error {
	if r, ok := TryWithErrorValue(callback); !ok {
		var err error
		if err, ok = r.(error); !ok {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		return err
	}

	return nil
}

// TryWithErrorValue has the same behavior as Try, but also returns value passed to panic.
// Play: https://go.dev/play/p/Kc7afQIT2Fs
func TryWithErrorValue(callback func() error) (errorValue any, ok bool) {
	ok = true

	defer func() {
		if r := recover(); r != nil {
			ok = false
			errorValue = r
		}
	}()

	err := callback()
	if err != nil {
		ok = false
		errorValue = err
	}

	return
}

func Retry(ctx context.Context, maxTimes int, interval time.Duration, f func() error) {
	for i := 0; i < maxTimes; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			err := f()
			if err == nil {
				return
			}
			time.Sleep(interval)
		}
	}
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}

func Md5Bytes(data []byte) string {
	h := md5.New()
	h.Write(data)
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}

func Base64Encode(data []byte) string {
	str := base64.StdEncoding.EncodeToString(data)
	str = strings.Replace(str, "+", "*", -1)
	str = strings.Replace(str, "/", "-", -1)
	str = strings.Replace(str, "=", "_", -1)
	return str
}

func Base64Decode(str string) ([]byte, error) {
	str = strings.Replace(str, "_", "=", -1)
	str = strings.Replace(str, "-", "/", -1)
	str = strings.Replace(str, "*", "+", -1)
	return base64.StdEncoding.DecodeString(str)
}
