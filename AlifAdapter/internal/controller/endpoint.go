package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(ctx *gin.Context) {

	route := gin.Default()
	route.POST("check", check)
	route.POST("pay", pay)
	route.POST("post_check", post_check)
	route.POST("account",account)

	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, "No request!!!")
	})

	route.Run()
}
func check(ctx *gin.Context) {

}

func pay(ctx *gin.Context) {

}

func post_check(ctx *gin.Context) {

}

func account(gin *gin.Context){
	
}
