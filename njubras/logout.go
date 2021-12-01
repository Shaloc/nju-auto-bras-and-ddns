package njubras

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const PortalLogoutURL = PortalBaseUrl + PortalLogout

type LogoutData struct {
	ReplyCode int `json:"reply_code"`
	ReplyMsg string `json:"reply_msg"`
}

func DoLogout() (LogoutData, error) {
	req, err := http.NewRequest("GET", PortalLogoutURL, nil)
	logoutData := LogoutData{}
	if err != nil {
		return logoutData, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return logoutData, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return logoutData, err
	}
	var data LogoutData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return logoutData, err
	}
	log.Println("BrasLogout: " + data.ReplyMsg)
	return logoutData, nil
}