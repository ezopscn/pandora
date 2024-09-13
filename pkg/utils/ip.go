package utils

import (
	"fmt"
	"net"
)

// 获取宿主机的内网 IP 地址
func GetHostIP() (string, error) {
	var ipAddresses []net.IP

	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	// 遍历每一个接口
	for _, iface := range interfaces {
		// 忽略无效的接口和回环接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 获取接口的所有地址
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		// 遍历每一个地址
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 检查是否是内网地址
			if ip != nil && ip.IsPrivate() {
				ipAddresses = append(ipAddresses, ip)
			}
		}
	}

	// 如果找到了内网 IP 地址，则返回第一个
	if len(ipAddresses) > 0 {
		return ipAddresses[0].String(), nil
	}

	return "", fmt.Errorf("没有获取到匹配的内网 IP 地址")
}
