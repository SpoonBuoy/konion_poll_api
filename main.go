package main

import (
	"poll/service"
	"poll/store"
)

func main() {
	dbStore, err := store.NewDatabase()
	cacheStore, err := store.NewCache()

	dbPollService, err := service.NewPollService(dbStore)
	cachePollServie, err := service.NewPollService(cacheStore)
}
