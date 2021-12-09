package main

import (
	"log"
	"nju_auto_ddns/ddns"
	"nju_auto_ddns/njubras"
	"nju_auto_ddns/notify"
	"time"
)

func App(conf *AutoDdnsConf, ddns ddns.IDynamicDNS, bark notify.INotifier) {
	// TODO: Support other ddns tools and notifiers
	for {
		_, err := njubras.DoLogin(conf.Bras.UserName, conf.Bras.Password)
		if err != nil {
			log.Println("App: BrasLogin failed with " + err.Error())
			continue
		}
		userData, err := njubras.AcquirePortalStatus()
		if err != nil {
			log.Println("App: Acquire bras status failed with " + err.Error())
			continue
		}
		if userData == nil || userData.Results.TotalElems == 0 {
			log.Println("App: Empty bras status response")
			continue
		}
		ipv4Addr := njubras.IntToIPv4(userData.Results.ResultElems[0].UserIPV4)
		log.Println("App: Current IP is " + ipv4Addr.String())
		ddns.Update(ipv4Addr)
		// TODO: determine where to notify, add a custom log system and add a trigger
		time.Sleep(time.Duration(conf.Core.UpdateInterval) * time.Second)
	}
}
