package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jessevdk/go-flags"
	"informator-69-bot/app/publisher"
	"informator-69-bot/app/storage"
	"informator-69-bot/app/wiki"
	"log"
)

var opts struct {
	ApiToken string `short:"t" long:"token" env:"API_TOKEN" description:"Telegram bot api token"`
	Storage  string `short:"s" long:"storage" env:"SUBSCR_STORAGE" description:"File to store subscribers"`
	Holidays string `short:"h" long:"holidays" env:"HOLIDAYS_FILE" description:"File to get holidays"`
}

func main() {
	if _, err := flags.Parse(&opts); err != nil {
		log.Panic(err)
	}

	store := storage.NewStore(opts.Storage)
	holidays := wiki.LoadHolidays(opts.Holidays)

	bot, err := tgbotapi.NewBotAPI(opts.ApiToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	go publisher.NotifierFromFile(store, &holidays, bot)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // ignore any non-Message Updates

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					go store.AddChat(update.Message.Chat.ID)
					publisher.SendMessage(update.Message.Chat.ID, "You will be receiving useful information", bot)
				case "stop":
					go store.RemoveChat(update.Message.Chat.ID)
					publisher.SendMessage(update.Message.Chat.ID, "You won't be receiving useful information", bot)
				case "info":
					publisher.SendMessage(update.Message.Chat.ID, wiki.GetTodaysReport(), bot)
				case "city":
					go store.AddCity(update.Message)
				case "full":
					// need to generate message and buttons
				}
			}
		} else if update.CallbackQuery != nil {
			log.Print("Get callback query")
		}
	}
}
