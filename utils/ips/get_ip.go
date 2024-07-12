package ips

import (
	"fmt"
	"net"
)

func GetIP() (Addr string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("获取网卡信息出错, err:", err)
		return
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("获取IP地址出错, err:", err)
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					fmt.Println("本机IP地址:", ipNet.IP.String())
					return ipNet.IP.String()
				}
			}
		}
	}
	return
}
