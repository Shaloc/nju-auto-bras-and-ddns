package provider

import (
	"context"
	"log"
	"net"
	"net/url"
	"nju_auto_ddns/ddns"
	"strings"

	"github.com/cloudflare/cloudflare-go"
)

type DDNSCloudflare struct {
	apiKey string
	apiEmail string
	domain string
	url string
	zoneID string
	lastIP net.IP
	instance *cloudflare.API
}

func (cf *DDNSCloudflare) Initialize(conf *ddns.DynamicDNSConfig) {
	tmpUrl, err := url.Parse("https://" + conf.TargetUrl)
	if err != nil {
		log.Fatalln("Cloudflare: Error: " + err.Error())
	}
	parts := strings.Split(tmpUrl.Hostname(), ".")
	domain := parts[len(parts) - 2] + "." + parts[len(parts) - 1]
	cf.apiKey = conf.ApiKey
	cf.apiEmail = conf.ApiEmail
	cf.url = conf.TargetUrl
	cf.domain = domain
	cf.lastIP = net.IP{0, 0, 0, 0}
	api, err := cloudflare.New(cf.apiKey, cf.apiEmail)
	if err != nil {
		log.Fatalln("Cloudflare: initialize api failed... " + err.Error())
	}
	cf.instance = api
	id, err := cf.instance.ZoneIDByName(cf.domain)
	if err != nil {
		log.Fatalln("Cloudflare: Error " + err.Error())
	}
	cf.zoneID = id
	log.Printf("Cloudlfare: Initialized as %#v", cf)
}


func (cf *DDNSCloudflare) Update(ip net.IP) {
	if ip.String() == cf.lastIP.String() {
		log.Println("Cloudflare: IP unchanged")
		return
	}
	log.Println("Cloudflare: " + cf.lastIP.String() + "->" + ip.String() + " for " + cf.url)
	copy(cf.lastIP, ip)
	originRecord := cloudflare.DNSRecord{
		Type:       "A",
		Name:       cf.url,
	}
	record := cloudflare.DNSRecord{
		Type:       "A",
		Name:       cf.url,
		Content:    ip.String(),
		TTL:        1,
	}

	records, err := cf.instance.DNSRecords(context.Background(), cf.zoneID, originRecord)
	if err != nil {
		log.Println("Cloudflare: Error: " + err.Error())
		return
	}
	if len(records) == 0 {
		rs, err := cf.instance.CreateDNSRecord(context.Background(), cf.zoneID, record)
		if err != nil {
			log.Println("Cloudflare: Error: " + err.Error())
			return
		}
		log.Printf("Cloudflare: Create new record %#v\n", rs)
		return
	}
	recordId := records[0].ID
	err = cf.instance.UpdateDNSRecord(context.Background(), cf.zoneID, recordId, record)
	if err != nil {
		log.Println("Cloudflare: Error: " + err.Error())
		return
	}
	log.Printf("Cloudflare: Updating record %#v\n", record)
}
