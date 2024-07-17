package main

import (
	"exp/endpoint"
	"exp/telegram"

	"github.com/gin-gonic/gin"
)



func main() {
	rout:=gin.Default()
	rout.GET("/ping",endpoint.CheckPing)
	rout.GET("/currency",endpoint.CheckCurrence)
	rout.GET("/tokensInfo",endpoint.TokensHistory)
	rout.GET("/binanceCon",endpoint.ConnectBinance)
	rout.GET("/currenciesList",endpoint.CurrenciesList)
	rout.Run(":3000")
	telegram.TelegramBOT()

}

