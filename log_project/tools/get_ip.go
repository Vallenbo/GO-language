package tools

import (
	"fmt"
	"net"
	"strings"
)

func GetOurboundIP() (ip string, err error) { // 获取本地对外IP
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}

func main() {
	ip, err := GetOurboundIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}
