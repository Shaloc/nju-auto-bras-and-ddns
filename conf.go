package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type AutoDdnsConf struct {
	Core struct {
		UpdateInterval int `yaml:"update_interval"`
	} `yaml:"core"`
	Provider struct {
		Using string `yaml:"using"`
		Url string `yaml:"url"`
		CloudFlare struct {
			Email string `yaml:"email"`
			ApiKey string `yaml:"api_key"`
		} `yaml:"cloudflare"`
	} `yaml:"provider"`
	Bras struct {
		UserName string `yaml:"user_name"`
		Password string `yaml:"password"`
	} `yaml:"bras"`
	Notify struct {
		Using string `yaml:"using"`
		Bark struct {
			Topic string `yaml:"topic"`
			Url string `yaml:"url"`
		}
	}
}

func ReadConfig(fileName string) (*AutoDdnsConf, error) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	conf := &AutoDdnsConf{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		log.Fatalf("ReadConf: Parse config failed %v", err)
	}
	log.Printf("ReadConfig: Load config success: %#v", conf)
	return conf, nil
}