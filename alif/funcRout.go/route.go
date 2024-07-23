package funcroutgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

)

func SendRequest(ids string) string {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/coins/markets"+"?vs_currency=usd&ids="+ids+"&category=layer-1", nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Set("x-cg-demo-api-key", "CG-zRqVxd1dzX5eiXpXYiu7YxPt")
	req.Header.Add("accept", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var value []ResTokensInfo
	err=json.Unmarshal(body,&value)
	if err!=nil{
		fmt.Println(err)
		return ""
	}
	return string(value[0].)
}