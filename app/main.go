package main

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"informator-69-bot/app/publisher"
	"informator-69-bot/app/storage"
	"informator-69-bot/app/wiki"
	"log"
	"os"
)

type Config struct {
	ApiToken string `json:"api_token"`
	Storage  string `json:"storage"`
}

func getConfig(CfgFileName string) (Config, error) {
	log.Println("Parse config file:", CfgFileName)
	file, err := os.Open(CfgFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Print(err)
		}
	}()
	decoder := json.NewDecoder(file)
	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		log.Fatal(err)
	}
	return config, nil
}

func main() {
	configFileName := os.Getenv("BOT_CONFIG")
	if configFileName == "" {
		configFileName = "config.json"
	}

	config, _ := getConfig(configFileName)
	store := storage.NewStore(config.Storage)

	bot, err := tgbotapi.NewBotAPI(config.ApiToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	go publisher.Notifier(store, bot)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				go store.Add(update.Message.Chat.ID)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You will be receiving useful information")
				if _, err := bot.Send(msg); err != nil {
					log.Print(err)
				}
			case "stop":
				go store.Remove(update.Message.Chat.ID)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You won't be receiving useful information")
				if _, err := bot.Send(msg); err != nil {
					log.Print(err)
				}
			case "info":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, wiki.GetTodaysReport())
				msg.ParseMode = "markdown"
				if _, err := bot.Send(msg); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
