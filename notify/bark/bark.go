package bark

import (
	"log"
	"net/http"
	"nju_auto_ddns/notify"
)

type NotifierBark struct {
	apiUrl string
	topic string
}

func (barkNotifier *NotifierBark) Initialize(conf *notify.NotifierConfig) {
	barkNotifier.apiUrl = conf.ApiUrl
	barkNotifier.topic = conf.Topic
}

func (barkNotifier *NotifierBark) DoNotify(message string) {
	if len(barkNotifier.apiUrl) == 0 {
		log.Println("BarkNotifier: Error: no api url set")
		return
	}
	queryUrl := barkNotifier.apiUrl + "/" + barkNotifier.topic + "/" + message
	log.Println("BarkNotifier: Debug: query " + queryUrl)
	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		log.Println("BarkNotifier: Error: Bark notify failed: " + err.Error())
		return
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println("BarkNotifier: Error: Bark notify failed: " + err.Error())
		return
	}
	log.Println("BarkNotifier: Sent " + barkNotifier.topic + " " + message)
	return
}