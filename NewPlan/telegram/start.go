package telegram

import (
	"encoding/json"
	"exp/endpoint"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func TelegramBOT() {
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
		case "balance":
			F:= func(ctx *gin.Context)any {
				ids := ctx.Query("ids")
				vs_currencies := ctx.Query("vs_currencies")
				include_24hr_change := ctx.Query("include_24hr_change")
				url := "https://api.coingecko.com/api/v3/simple/price" + "?ids=" + ids + "&vs_currencies=" + vs_currencies + "&include_24hr_change=" + include_24hr_change

				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					logrus.WithError(err).Info("on request in CheckCurrence")
					return err
				}

				req.Header.Add("accept", "application/json")
				req.Header.Add("x-cg-demo-api-key", "CG-zRqVxd1dzX5eiXpXYiu7YxPt\t")

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					logrus.WithError(err).Info("on respoonse in CheckCurrence")
					return err
				}

				defer res.Body.Close()
				body, err := io.ReadAll(res.Body)
				if err != nil {
					logrus.WithError(err).Info("on reading body in CheckAccount")
					return err
				}

				var value endpoint.ABS
				err = json.Unmarshal(body, &value)
				if err != nil {
					logrus.WithError(err).Info("on marshaling in CheckCurrency")
					return err
				}
				fmt.Println(string(body))
				ctx.IndentedJSON(200, value)
				return value
			}
			fmt.Println("hello world")
			fmt.Println(F)
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
