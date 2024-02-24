package utils

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"net/http"
)

func GetHttpRemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

func SplitHostPort(addr string) (host, port string, err error) {
	return net.SplitHostPort(addr)
}

func Ip2long(s string) uint32 {
	ip := net.ParseIP(s)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func GetIdlePort() string {
	var a *net.TCPAddr
	var err error
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return fmt.Sprintf("0.0.0.0:%d", l.Addr().(*net.TCPAddr).Port)
		}
	}
	log.Fatalf("get rpc port failed: %v\n", err)
	return ""
}
