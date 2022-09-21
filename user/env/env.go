package env

import (
	"net"
	"os"
)

var (
	AppName  = "order"
	RootPath = ""
	ConfPath = ""
	LogPath  = ""
	ContextKeyRequestID = "requestId"
	TraceHeaderKey      = "Uber-Trace-Id"
	UriNow              = "uriNow"
)

func GetLocalIp() string {
	addrs, _ := net.InterfaceAddrs()
	var ip string = ""
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				if ip != "127.0.0.1" {
					return ip
				}
			}
		}
	}
	return "127.0.0.1"
}

func InitPath() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	RootPath = path + "/"
	ConfPath = RootPath + "conf/"
	LogPath = RootPath + "log/"
}
