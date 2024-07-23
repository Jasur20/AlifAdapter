package main

import (
	funcroutgo "exp/funcRout.go"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// func main(){
// 	a:=1
// 	if a==1{
// 		c := http.Client{Timeout: time.Duration(1) * time.Second}
// 		req,err:=http.NewRequest("GET","https://api.coingecko.com/api/v3/simple/price"+"?ids=bitcoin&vs_currencies=usd",nil)
// 		if err!=nil{
// 			fmt.Println(err)
// 			return
// 		}

// 		req.Header.Set("x-cg-demo-api-key","CG-zRqVxd1dzX5eiXpXYiu7YxPt")
// 		req.Header.Add("accept", "application/json")

// 		resp, err :=c.Do(req)
// 		if err!=nil{
// 			fmt.Println(err)
// 			return
// 		}

// 		defer resp.Body.Close()

// 		body,err:=ioutil.ReadAll(resp.Body)
// 		if err!=nil{
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Print(string(body))
// 		return

//		}
//		fmt.Println("no 1")
//	}

type ResCoin map[string]map[string]any

func main() {
	bot, err := tgbotapi.NewBotAPI("6591336579:AAGqkqnAeCYn9SBSsPDJXW7N_asWkgDFb5k")
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "hi":
			resp := funcroutgo.SendRequest("bitcoin")
			msg.Text = string(resp)
		// case "eth":
		// 	resp:=funcroutgo.SendRequest("ethereum")
		// 	msg.Text=string(resp)
		// case "not":
		// 	resp:=funcroutgo.SendRequest("notcoin")
		// 	msg.Text=string(resp)
		// case "sol":
		// 	resp:=funcroutgo.SendRequest("solana")
		// 	msg.Text=string(resp)
		// case "ton":
		// 	resp:=funcroutgo.SendRequest("the-open-network")
		// 	msg.Text=string(resp)
		// case "near":
		// 	resp:=funcroutgo.SendRequest("near")
		// 	msg.Text=string(resp)
		// case "pepe":
		// 	resp:=funcroutgo.SendRequest("pepe")
		// 	msg.Text=string(resp)
		default:
			msg.Text = "I don't know that command"
		}
		fmt.Println("hello world")
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
