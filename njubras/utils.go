package njubras

import (
	"encoding/binary"
	"net"
)

const (
	PortalBaseUrl = "http://p.nju.edu.cn/"
	PortalLogin = "portal_io/login"
	PortalLogout = "portal_io/logout"
	PortalStatus = "api/portal/v1/getinfo"
)

func IPv4ToInt(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func IntToIPv4(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}