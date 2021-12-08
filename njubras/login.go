package njubras

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const PortalLoginURL = PortalBaseUrl + PortalLogin

type UserInfo struct {
	UserName string `json:"username"`
	FullName string `json:"fullname"`
	ServiceName string `json:"service_name"`
	AreaName string `json:"area_name"`
	AcctStartTime uint `json:"acctstarttime"`
	Balance uint `json:"balance"`
	UserIPV4 uint32 `json:"useripv4"`
	UserIPV6 string `json:"useripv6"`
	MacAddr string `json:"mac"`
	Domain string `json:"domain"`
	PortalServerIP uint `json:"portal_server_ip"`
	PortalAcctSessionID uint `json:"portal_acctsessionid"`
}

type LoginData struct {
	ReplyCode int `json:"reply_code"`
	ReplyMsg string `json:"reply_msg"`
	Info UserInfo `json:"user_info"`
	RequestURI string `json:"request_uri"`
	RequestTime uint `json:"request_time"`
}

func DoLogin(UserName, Password string) (LoginData, error) {
	fullURL := PortalLoginURL + "?username=" + UserName + "&password=" + Password
	loginData := LoginData{}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return loginData, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return loginData, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return loginData, err
	}
	err = json.Unmarshal(body, &loginData)
	if err != nil {
		return loginData, err
	}
	log.Println("BrasLogin: " + loginData.ReplyMsg)
	return loginData, nil
}