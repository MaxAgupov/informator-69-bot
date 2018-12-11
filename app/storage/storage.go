package storage

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Chat struct {
	Id int64 `json:"id"`
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

func (store *ActiveChatsStore) Add(Id int64) {
	newChat := Chat{Id}
	store.Lock()
	defer store.Unlock()
	store.Cache[Id] = newChat
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

func (store *ActiveChatsStore) Remove(Id int64) {
	store.Lock()
	defer store.Unlock()
	delete(store.Cache, Id)
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

