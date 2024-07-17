package endpoint

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// func SendForGet(ctx *gin.Context){

// 	resp, err := http.Get("https://evilinsult.com/generate_insult.php?lang=en&type=json")
// 	if err!=nil{
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	body,err:=io.ReadAll(resp.Body)
// 	if err!=nil{
// 		fmt.Println(err)
// 		return
// 	}
// var value *Take
// err=json.Unmarshal(body,&value)
// if err!=nil{
// 	ctx.JSON(404,"error message")
// 	return
// }

// fmt.Println(value.Insult)
// ctx.JSON(200,gin.H{
// 	"message":value,
// }


func CheckPing(ctx *gin.Context){
	
	url := "https://api.coingecko.com/api/v3/ping"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-demo-api-key", "CG-zRqVxd1dzX5eiXpXYiu7YxPt\t")

	res, err := http.DefaultClient.Do(req)
	if err!=nil{
		logrus.WithError(err).Warn("on response in ChackPing")
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err!=nil{
		logrus.WithError(err).Info("on body in CheckPing")
		return
	}

	var value *Ping
	err=json.Unmarshal(body,&value)
	if err!=nil{
		logrus.WithError(err).Info("on marshalling in CheckPing")
		return
	}

	ctx.IndentedJSON(200,value.GeckoSays)

}

func CheckCurrence(ctx *gin.Context){

	ids:=ctx.Query("ids")
    vs_currencies:=ctx.Query("vs_currencies")
	include_24hr_change:=ctx.Query("include_24hr_change")
	url := "https://api.coingecko.com/api/v3/simple/price"+"?ids="+ids+"&vs_currencies="+vs_currencies+"&include_24hr_change="+include_24hr_change

	req, err:= http.NewRequest("GET", url, nil)
	if err!=nil{
		logrus.WithError(err).Info("on request in CheckCurrence")
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-demo-api-key", "CG-zRqVxd1dzX5eiXpXYiu7YxPt\t")

	res, err := http.DefaultClient.Do(req)
	if err!=nil{
		logrus.WithError(err).Info("on respoonse in CheckCurrence")
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err!=nil{
		logrus.WithError(err).Info("on reading body in CheckAccount")
		return
	}


	var value ABS
	err=json.Unmarshal(body,&value)
	if err!=nil{
		logrus.WithError(err).Info("on marshaling in CheckCurrency")
		return
	}
	fmt.Println(string(body))

	ctx.IndentedJSON(200,value)
}

func TokensHistory(ctx *gin.Context){
	
	vs_currency:=ctx.Query("vs_currency")
	// id:=ctx.Query("id")
	ids:=ctx.Query("ids")
	category:=ctx.Query("category")
	url:="https://api.coingecko.com/api/v3/coins/markets"+"?vs_currency="+vs_currency+/*"&id="+id+*/"&ids="+ids+"&category="+category

	req, err:= http.NewRequest("GET", url, nil)
	if err!=nil{
		logrus.WithError(err).Info("on request in TokensHistory")
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-demo-api-key", "CG-zRqVxd1dzX5eiXpXYiu7YxPt\t")

	res, err := http.DefaultClient.Do(req)
	if err!=nil{
		logrus.WithError(err).Info("on respoonse in TokensHistory")
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err!=nil{
		logrus.WithError(err).Info("on reading body in TokensHistory")
		return
	}
	fmt.Println(string(body))

	
	var value []ResTokensInfo
	err=json.Unmarshal(body,&value)
	if err!=nil{
		logrus.WithError(err).Info("on marshaling in TokensHistory")
		return
	}
	for i:=range(value){
		if value[i].Symbol=="btc"{
			ctx.IndentedJSON(http.StatusOK,value[i].Symbol)
			return
		}
		
	}
}

func CurrenciesList(ctx *gin.Context){
	
	url := "https://api.coingecko.com/api/v3/simple/supported_vs_currencies"

	req, err:= http.NewRequest("GET", url, nil)
	if err!=nil{
		logrus.WithError(err).Info("on request in CurrenciesList")
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-demo-api-key", "CG-zRqVxd1dzX5eiXpXYiu7YxPt\t")

	res, err := http.DefaultClient.Do(req)
	if err!=nil{
		logrus.WithError(err).Info("on respoonse in CurrenciesList")
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err!=nil{
		logrus.WithError(err).Info("on reading body in CurrenciesList")
		return
	}
	fmt.Println(string(body))
	ctx.IndentedJSON(http.StatusOK,string(body))
}

func ConnectBinance(ctx *gin.Context){
}

