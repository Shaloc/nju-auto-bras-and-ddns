package ddns

import "net"

type DynamicDNSConfig struct {
	ApiKey string
	ApiEmail string
	TargetUrl string
}

type IDynamicDNS interface {
	Initialize(data *DynamicDNSConfig)
	Update(addr net.IP)
}
