package main

import (
	"flag"
	"log"
	"nju_auto_ddns/ddns"
	"nju_auto_ddns/ddns/provider"
	"nju_auto_ddns/notify"
	bark2 "nju_auto_ddns/notify/bark"
	"nju_auto_ddns/notify/placeholder"
)

func main() {
	confFile := flag.String("c", "template.yaml", "Your config file")
	flag.Parse()
	conf, err := ReadConfig(*confFile)
	if err != nil {
		log.Fatal(err)
	}
	var dynamicDNS ddns.IDynamicDNS
	var notifier notify.INotifier

	switch conf.Provider.Using {
	case "":
		dynamicDNS = &provider.DDNSNil{}
		log.Println("App: DDNS feature will not be used.")
	case "cloudflare":
		dynamicDNS = &provider.DDNSCloudflare{}
		cfCfg := &ddns.DynamicDNSConfig{
			ApiKey: conf.Provider.CloudFlare.ApiKey,
			ApiEmail: conf.Provider.CloudFlare.Email,
			TargetUrl: conf.Provider.Url}
		dynamicDNS.Initialize(cfCfg)
	default:
		log.Fatalln("DDNS provider " + conf.Provider.Using + " has not yet been implemented." )
	}

	switch conf.Notify.Using {
	case "":
		notifier = &placeholder.NotifierNil{}
		log.Println("App: Notify feature will not be used.")
	case "bark":
		notifier = &bark2.NotifierBark{}
		notifier.Initialize(&notify.NotifierConfig{ApiUrl: conf.Notify.Bark.Url, Topic: conf.Notify.Bark.Topic})
	default:
		log.Fatalln("Notifier " + conf.Notify.Using + " has not yet been implemented.")
	}

	App(conf, dynamicDNS, notifier)
}
