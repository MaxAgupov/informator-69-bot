package storage

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"informator-69-bot/app/weather"
	"log"
	"os"
	"sync"
)

type Chat struct {
	Id int64 `json:"id"`
	//Cities *[]City `json:"cities"`
}

type City struct {
	UserId   int    `json:"user_id"`
	CityId   int64  `json:"city_id"`
	CityName string `json:"city_name"`
}

type ActiveChatsStore struct {
	sync.RWMutex
	storage string
	Cache   map[int64]Chat
}

func NewStore(Storage string) *ActiveChatsStore {
	log.Println("Load data from storage:", Storage)
	file, err := os.OpenFile(Storage, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Print(err)
		}
	}()
	decoder := json.NewDecoder(file)
	activeChats := make(map[int64]Chat)
	if err := decoder.Decode(&activeChats); err != nil {
		log.Print(err)
	}
	return &ActiveChatsStore{
		storage: Storage,
		Cache:   activeChats,
	}
}

func (store *ActiveChatsStore) AddChat(ChatId int64) {
	//newChat := Chat{ChatId, nil}
	newChat := Chat{ChatId}
	store.Lock()
	defer store.Unlock()
	store.Cache[ChatId] = newChat
	if err := os.Rename(store.storage, store.storage+".bak"); err != nil {
		log.Println("Can't create storage backup:", err)
		return
	}
	file, _ := os.Create(store.storage)
	defer func() {
		if err := file.Close(); err != nil {
			log.Print(err)
		}
	}()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(store.Cache); err != nil {
		log.Print(err)
	}
}

func (store *ActiveChatsStore) RemoveChat(ChatId int64) {
	store.Lock()
	defer store.Unlock()
	delete(store.Cache, ChatId)
	if err := os.Rename(store.storage, store.storage+".bak"); err != nil {
		log.Println("Can't create storage backup:", err)
		return
	}
	file, _ := os.Create(store.storage)
	defer func() {
		if err := file.Close(); err != nil {
			log.Print(err)
		}
	}()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(store.Cache); err != nil {
		log.Print(err)
	}
}

func (store *ActiveChatsStore) AddCity(Msg *tgbotapi.Message) {
	log.Println("Msg.Chat.ID=", Msg.Chat.ID)
	log.Println("Msg.Chat.Type=", Msg.Chat.Type)
	if Msg.From != nil {
		log.Println("Msg.From.ID=", Msg.From.ID)
		log.Println("Msg.From.ID=", Msg.From.UserName)
		cityName := Msg.CommandArguments()
		city := NewCity(Msg.From.ID, cityName)
		log.Println("City validation = ", city.ValidateCity())
	} else {
		log.Println("Msg.From=", Msg.From)
	}
	log.Println("Msg=", Msg.CommandArguments())
}

func NewCity(UserId int, CityNameRu string) City {
	return City{UserId, 0, CityNameRu}
}

func (city *City) ValidateCity() bool {
	weather := weather.GetCurrentWeather(city.CityName)
	if weather != nil {
		return true
	}
	return false
}
