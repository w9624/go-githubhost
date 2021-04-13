package runner

import (
	"fmt"
	"net"
)

func ParseAddr(domains []string) map[string]string {
	addrMap := make(map[string]string, len(domains))

	for _, domain := range domains {
		addr, err := net.ResolveIPAddr("ip", domain)
		if err == nil {
			addrMap[domain] = addr.IP.String()
		} else {
			addrMap[domain] = fmt.Sprintf("解析异常：%s", err.Error())
		}
	}

	return addrMap
}
