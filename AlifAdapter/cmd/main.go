package main

import (
	"alif/internal/adapter"
	"alif/internal/controller"
	"alif/internal/integration"
	"fmt"
	"net"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


type Config struct{
	EndPoint EndPoint `mapstructure:"endpoint"`
	Adapter Adapter `mapstructure:"adapter"`
}

type EndPoint struct{
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Token string `mapstructure:"token"`
}

type Adapter struct{
	Url string `mapstructure:"url"`
	UserID string `mapstructure:"userid"`
	Password string `mapstructure:"password"`
	TimeOut int `mapstructure:"timeout"`
}


var Settings Config

func ReadConfig() Config{
	viper.SetConfigName("config") 
	viper.SetConfigType("yml")   
	viper.AddConfigPath(".")      
	
	if err := viper.ReadInConfig(); err != nil { 
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return config
}



func customClient(timeout time.Duration) *http.Client {
	
	transport := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   1 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	client := http.Client{
		Transport: &transport,
		Timeout:   timeout * time.Second,
	}

	return &client
}


func main() {
	var config Config
	config=ReadConfig()
	httpClient:=customClient(time.Duration(config.Adapter.TimeOut))
	adapter:=adapter.NewAdapter(config.Adapter.Url,config.Adapter.UserID,httpClient,config.Adapter.Password)
	integration:=integration.NewIntegration(adapter)
	
	engine:=gin.Default()
	endpoint:=controller.NewEndpointController(config.EndPoint.Token,engine,integration)
	endpoint.InitRoutes()
	engine.Run(config.EndPoint.Host+":"+fmt.Sprint(config.EndPoint.Port))
}