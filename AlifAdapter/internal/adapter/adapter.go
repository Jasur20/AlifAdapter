package adapter

import (
	"alif/internal/integration"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/sirupsen/logrus"
)


type adapter struct {
	cftUrl        string
	alifUrl       string
	userid        string
	httpClient    *http.Client
}

func NewAdapter(cftUrl string, alifUrl ,userid string,httpClient *http.Client) integration.Adapter {
	return &adapter{cftUrl: cftUrl,alifUrl: alifUrl, userid: userid, httpClient: httpClient}
}

func (a adapter) PreCheck(account string, serviceID string) (status int64, description string, rawInfo map[string]string, err error) {

	if slices.Contains(a.pans, account[0:9]) {

		status, description, rawInfo, err = cftCheck(account, a.cftUrl, a.httpClient)

	} else {
        var resAlif Resp
		resAlif,err=alifCheck(account,serviceID,a.alifUrl,a.userid,a.httpClient)
	}

	return status, description, rawInfo, err
}

func alifCheck(account,serviceID,alifUrl,userid string,httpClien *http.Client) (Resp,error){

    u,err:=json.Marshal(Req{Service:serviceID,UserID: userid,Hash: "",Account: account,Amount: "",Currency: "TJS",TxnID: "",Phone: "",
	Fee: "",Providerid: "0",Last_Name: "",First_Name: "",Middle_Name: "",Sender_Birthday: "",Address: "",Resident_Country: "",Postal_Code: "",Recipient_Name: ""})
	if err!=nil{
		logrus.WithError(err).Warn("on marshaling")
		return Resp{},err
	}
	// encodedBody := data.Encode()
	req,err:=http.NewRequest(http.MethodGet,alifUrl,bytes.NewBuffer(u))
	if err!=nil{
		logrus.WithError(err).Warn("on create request")
		return Resp{},nil
	}
	req.Header.Add("Content-Type", "application/json; charset=utf8")

	if err!=nil{
		logrus.WithError(err).Warn("on creating request")
		return Resp{},nil
	}
	res,err:=httpClien.Do(req)
	if err!=nil{
		logrus.WithError(err).Info("on doing request")
		return Resp{},nil
	}
	defer res.Body.Close()

	body,err:=io.ReadAll(res.Body)
	if err!=nil{
		logrus.WithError(err).Warn("on reading responce body")
		return Resp{},nil
	}
	
	logrus.WithFields(logrus.Fields{
		"body": string(body),
	}).Info("response to adapter")

	resBody:=&Resp{}
	err=json.Unmarshal(body,resBody)
	if err!=nil{
		logrus.WithError(err).Info("on parsing responce")
		return Resp{},nil
	}

	return *resBody,nil
}


func cftCheck(account string, url string, httpClient *http.Client) (status int64, description string, rawInfo map[string]string, err error) {
	req, err := http.NewRequest(http.MethodGet, url+"/card/holder/"+account, nil)
	if err != nil {
		logrus.WithError(err).Info("on creating request")
		return 0, "", nil, err
	}
	res, err := httpClient.Do(req)

	if err != nil {
		logrus.WithError(err).Warn("on doing request")
		return 0, "", nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.WithError(err).Warn("on reading response body")
		return 0, "", nil, err
	}

	logrus.WithFields(logrus.Fields{
		"body": string(body),
	}).Info("response to adapter")
	resBody := &PreCheckResp{}
	err = json.Unmarshal(body, resBody)
	if err != nil {
		logrus.WithError(err).Info("on parsing response")
		return 0, "", nil, err
	}
	rawInfo = make(map[string]string)
	rawInfo["message"] = resBody.Message
	//rawInfo["currency"] = resBody.Currency

	return int64(res.StatusCode), res.Status, rawInfo, nil
}