package provider

import (
	"net"
	"nju_auto_ddns/ddns"
)

type DDNSNil struct {
}

func (ddnsNil *DDNSNil) Initialize(_ *ddns.DynamicDNSConfig) {
}

func (ddnsNil *DDNSNil) Update(_ net.IP) {
}