package provider

import (
	"net"
	"testing"
	"time"
)

func TestDDNSCloudflare_Initialize(t *testing.T) {
	cf := DDNSCloudflare{}
	// enter your cloudflare key, email and url you want to change here
	cf.Initialize("", "", "")
	ip := net.IP{114, 212, 114, 212}
	cf.Update(ip)
	time.Sleep(1000 * time.Millisecond)
	ip[3] = 213
	cf.Update(ip)
	time.Sleep(1000 * time.Millisecond)
	cf.Update(ip)
}
