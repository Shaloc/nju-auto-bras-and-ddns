package main

import (
	"flag"
	"log"
	bark2 "nju_auto_ddns/notify/bark"
	"nju_auto_ddns/provider"
)

func main() {
	confFile := flag.String("c", "template.yaml", "your config file")
	flag.Parse()
	conf, err := ReadConfig(*confFile)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: adapt other ddns providers
	cf := &provider.DDNSCloudflare{}
	cf.Initialize(conf.Provider.CloudFlare.ApiKey, conf.Provider.CloudFlare.Email, conf.Provider.Url)

	// TODO: adapt other notify tools
	bark := &bark2.NotifierBark{}
	bark.Initialize(conf.Notify.Bark.Url)
	App(conf, cf, bark)
}
