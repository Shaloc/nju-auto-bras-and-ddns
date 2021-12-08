package main

import (
	"log"
	"testing"
)

func TestReadConfig(t *testing.T) {
	conf, err := ReadConfig("template.yaml")
	if err != nil {
		panic(err)
	}
	log.Printf("%v", conf)
	log.Println(conf.Bras.UserName + " " + conf.Bras.Password)
	log.Println(conf.Provider.Using + " " + conf.Provider.CloudFlare.ApiKey)
	log.Println(conf.Notify.Using)
}
