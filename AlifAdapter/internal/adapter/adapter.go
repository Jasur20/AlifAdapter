package adapter

import (
	"alif/internal/integration"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)


type adapter struct {
	alifUrl       string
	userid        string
	httpClient    *http.Client
}

func NewAdapter(alifUrl string ,userid string,httpClient *http.Client) integration.Adapter {
	return &adapter{alifUrl: alifUrl, userid: userid, httpClient: httpClient}
}

func GetSha256(text string, secret []byte) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, secret)

	// Write Data to it
	h.Write([]byte(text))

	// Get result and encode as hexadecimal string
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}

func (a adapter) PreCheck(account string, serviceID string) (status int64, description string, rawInfo map[string]string, err error) {

	hash:=GetSha256(a.userid+account+"0"+"0",[]byte(a.userid))
    data,err:=json.Marshal(Req{Service:serviceID,UserID: a.userid,Hash: hash,Account: account,Amount: "",Currency: "TJS",TxnID: "",Phone: "",
	Fee: "",Providerid: "0",Last_Name: "",First_Name: "",Middle_Name: "",Sender_Birthday: "",Address: "",Resident_Country: "",Postal_Code: "",Recipient_Name: ""})
	if err!=nil{
		logrus.WithError(err).Warn("on marshaling")
		return 0,"0",rawInfo,nil
	}
	// encodedBody := data.Encode()
	req,err:=http.NewRequest(http.MethodGet,a.alifUrl,bytes.NewBuffer(data))
	if err!=nil{
		logrus.WithError(err).Warn("on create request")
		return 0,"0",rawInfo,nil
	}
	req.Header.Add("Content-Type", "application/json; charset=utf8")

	if err!=nil{
		logrus.WithError(err).Warn("on creating request")
		return 0,"0",rawInfo,nil
	}
	res,err:=a.httpClient.Do(req)
	if err!=nil{
		logrus.WithError(err).Info("on doing request")
		return 0,"0",rawInfo,nil
	}
	defer res.Body.Close()

	body,err:=io.ReadAll(res.Body)
	if err!=nil{
		logrus.WithError(err).Warn("on reading responce body")
		return 0,"0",rawInfo,nil
	}
	
	logrus.WithFields(logrus.Fields{
		"body": string(body),
	}).Info("response to adapter")

	resBody:=Resp{}
	err=json.Unmarshal(body,&resBody)
	if err!=nil{
		logrus.WithError(err).Info("on parsing responce")
		return 0,"0",rawInfo,nil
	}
	rawInfo = make(map[string]string)
	rawInfo["message"] = resBody.Message
	return int64(resBody.Code),resBody.Status,rawInfo,nil

}

func (a adapter) Payment(account string,serviceID string,amount string,trnID string, notifyRoute string) (status int64, description string, paymentID string, err error){

	hash:=GetSha256(a.userid+account+trnID+"0",[]byte(a.userid))
	temp,err:=json.Marshal(Req{Service:serviceID,UserID: a.userid,Hash: hash,Account: account,Amount: amount,Currency: "TJS",TxnID: "",Phone: "",
	Fee: "",Providerid: "0",Last_Name: "",First_Name: "",Middle_Name: "",Sender_Birthday: "",Address: "",Resident_Country: "",Postal_Code: "",Recipient_Name: ""})
	if err!=nil{
		logrus.WithError(err).Warn("on marshaling")
		return 0,"","0",err 
	}
	req,err:=http.NewRequest(http.MethodPost,a.alifUrl,bytes.NewBuffer(temp))
	req.Header.Add("Content-Type","application/json; charset=utf-8")

	if err!=nil{
		logrus.WithError(err).Info("on creating request")
	}
	res,err:=a.httpClient.Do(req)
	if err!=nil{
		logrus.WithError(err).Info("on doing request")
		return 0,"","0",err
	}
	defer res.Body.Close()

	body,err:=io.ReadAll(res.Body)
	if err!=nil{
		logrus.WithError(err).Info("on reading response body")
		return 0,"","0",err
	}

	logrus.WithFields(logrus.Fields{
		"body": string(body),
	}).Info("response to adapter")
	
	resBody:=&PaymentRes{}
	err=json.Unmarshal(body,resBody)
	if err!=nil{
		logrus.WithError(err).Info("on parsing response")
		return 0, "","0",err
	}
	return int64(resBody.StatusCode),resBody.Message,string(resBody.ID),nil
	
}

func (a adapter) PostCheck(trnID string,account string,serviceID string) (status int64, description string,err error){
	hash:=GetSha256(a.userid+account+trnID+"0",[]byte(a.userid))

	data:=Req{Service:serviceID,UserID: a.userid,Hash: hash,Account: account,Amount: "",Currency: "TJS",TxnID: trnID,Phone: "",
	Fee: "",Providerid: "0",Last_Name: "",First_Name: "",Middle_Name: "",Sender_Birthday: "",Address: "",Resident_Country: "",Postal_Code: "",Recipient_Name: ""}
	if err!=nil{
		logrus.WithError(err).Info("on creating data")
		return 0,"",nil
	}
	reqbody,err:=json.Marshal(data)
	if err!=nil{
		logrus.WithError(err).Info("on marshaling data")
		return 0,"",nil
	}
	req,err:=http.NewRequest(http.MethodPost,a.alifUrl,bytes.NewBuffer(reqbody))
	if err!=nil{
		logrus.WithError(err).Info("on req")
		return 0,"",nil
	}
	res,err:=a.httpClient.Do(req)
	if err!=nil{
		logrus.WithError(err).Info("on response")
		return 0,"",nil
	}
	defer res.Body.Close()
	body,err:=io.ReadAll(res.Body)
	if err!=nil{
		logrus.WithError(err).Info("on reading responce body")
		return 0,"",nil
	}
	logrus.WithFields(
		logrus.Fields{
			"body": string(body),
			"status": res.StatusCode,
		},
	).Info("response to adapter")

	resBody:=Resp{}
	err=json.Unmarshal(body,&resBody)
	if err!=nil{
		logrus.WithError(err).Info("on unmarshaling resBoby")
		return 0,"0",nil
	}
	return int64(resBody.Code),resBody.Message,nil
}	

