package njubras

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const PortalStatusUrl = PortalBaseUrl + PortalStatus

type ResultDataElem struct {
	AcctSessionID string `json:"actsessionid"`
	AcctStartTime int `json:"acctstarttime"`
	AreaName string `json:"area_name"`
	Balance int `json:"balance"`
	Domain string `json:"domain"`
	FullName string `json:"fullname"` // the actual name, like 张三
	MAC string `json:"mac"`
	ServiceName string `json:"service_name"`
	UserIPV4 int `json:"user_ipv4"`
	UserIPV6 string `json:"user_ipv6"`
	UserName string `json:"username"` // the NJUID

}

type ResultData struct {
	ResultElems []ResultDataElem `json:"rows"`
	TotalElems int `json:"total"`
}

type StatusData struct {
	ReplyCode int `json:"reply_code"`
	Results ResultData `json:"results"`
	ServerTime int `json:"server_time"`
}

// AcquirePortalStatus
// Log the current portal status, for debug purpose
// Handles the not login status
func AcquirePortalStatus() error {
	req, err := http.NewRequest("GET", PortalStatusUrl, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var data StatusData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	if data.ReplyCode != 0 {
		log.Println("BrasStatus: Cannot acquire bras status, maybe not login?" )
		log.Printf("BrasStatus: Error code %d\n", data.ReplyCode)
		return nil
	}
	for _, resultElem := range data.Results.ResultElems {
		log.Printf("BrasStatus: %v\n", resultElem)
		log.Printf("BrasStatus: Current IP4 %v\n", IntToIPv4(uint32(resultElem.UserIPV4)))
	}
	return nil
}