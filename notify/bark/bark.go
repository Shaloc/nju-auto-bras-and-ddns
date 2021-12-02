package bark

import (
	"log"
	"net/http"
)

type NotifierBark struct {
	apiUrl string
}

func (barkNotifier *NotifierBark) GetApiUrl() string {
	return barkNotifier.apiUrl
}

func (barkNotifier *NotifierBark) Initialize(url string) {
	barkNotifier.apiUrl = url
}

func (barkNotifier *NotifierBark) DoNotify(topic string, message string) {
	if len(barkNotifier.apiUrl) == 0 {
		log.Println("BarkNotifier: Error: no api url set")
		return
	}
	queryUrl := barkNotifier.apiUrl + "/" + topic + "/" + message
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
	log.Println("BarkNotifier: Sent " + topic + " " + message)
	return
}