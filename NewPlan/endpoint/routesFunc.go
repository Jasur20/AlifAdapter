package endpoint


import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)


func SendForGet(ctx *gin.Context){
	
	resp, err := http.Get("https://evilinsult.com/generate_insult.php?lang=en&type=json") 
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()
	body,err:=io.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println(err)
		return
	}
	var value *Take
	err=json.Unmarshal(body,&value)
	if err!=nil{
		ctx.JSON(404,"error message")
		return 
	}

	fmt.Println(value.Insult)
	ctx.JSON(200,gin.H{
		"message":value,
	})
}

func SendSecondreq(ctx *gin.Context){

	ids:=ctx.Param("ids")
	vs_currencies:=ctx.Param("vs_currencies")


	url := "https://api.coingecko.com/api/v3/simple/price"+"?ids="+ids+"&vs_currencies="+vs_currencies

	req, err:= http.NewRequest("GET", url, nil)
	if err!=nil{
		ctx.IndentedJSON(400,gin.H{
			"message":err,
		})
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-demo-api-key", "CG-2WBdT6jpFtQCk29mZ8i98y6N\t")

	res, err := http.DefaultClient.Do(req)
	if err!=nil{
		ctx.IndentedJSON(400,gin.H{
			"message":err,
		})
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err!=nil{
		ctx.IndentedJSON(400,gin.H{
			"message":err,
		})
		return
	}

	ctx.IndentedJSON(200,string(body))
}
