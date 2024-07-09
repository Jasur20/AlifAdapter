package main

import (
	"exp/endpoint"
	"github.com/gin-gonic/gin"
)



func main() {
	rout:=gin.Default()
	rout.GET("/balance",endpoint.SendSecondreq)
	rout.GET("/currency",)
	rout.Run(":3000")
}

