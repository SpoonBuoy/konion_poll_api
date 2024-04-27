package main

import (
	"fmt"
	"log"
	"poll/controller"
	"poll/service"
	"poll/store"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASWD    string
	DB_NAME     string
	REDIS_ADDR  string
	REDIS_PASWD string
	REDIS_DB    string
)

func init() {
	//set paths
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading config file %v", err.Error())
	}

	//config file
	PORT = fmt.Sprintf(":%v", viper.Get("server.port"))
	DB_HOST = fmt.Sprintf(":%v", viper.Get("db.host"))
	DB_PORT = fmt.Sprintf(":%v", viper.Get("db.port"))
	DB_NAME = fmt.Sprintf(":%v", viper.Get("db.name"))
	DB_USER = fmt.Sprintf(":%v", viper.Get("db.user"))
	DB_PASWD = fmt.Sprintf(":%v", viper.Get("db.password"))
	REDIS_ADDR = fmt.Sprintf(":%v", viper.Get("redis.addr"))
	REDIS_PASWD = fmt.Sprintf(":%v", viper.Get("redis.password"))
	REDIS_DB = fmt.Sprintf(":%v", viper.Get("redis.db"))
	log.Printf("CONFIG \n Port %s \n DB : %s \n Redis %s", PORT, DB_HOST+DB_PORT, REDIS_ADDR)

}

func main() {
	r := gin.Default()

	dbStore, err := store.NewDatabase()
	if err != nil {
		log.Printf("Error %v", err.Error())
	}
	cacheStore, err := store.NewCache()
	if err != nil {
		log.Printf("Error %v", err.Error())
	}

	dbPollService, err := service.NewPollService(dbStore)
	if err != nil {
		log.Printf("Error %v", err.Error())
	}
	cachePollService, err := service.NewPollService(cacheStore)

	if err != nil {
		log.Printf("Error %v", err.Error())
	}

	dbPollController := controller.NewPollController(dbPollService)
	cachePollController := controller.NewPollController(cachePollService)

	api := r.Group("/api")
	{
		api.GET("/poll/:{id}", dbPollController.FetchPoll)
		api.POST("/poll", dbPollController.CreatePoll)
		api.POST("/member", dbPollController.CreateMember)
		api.POST("/vote", cachePollController.RegisterVote)
		api.POST("/ref", cachePollController.CreateReference)
		api.POST("/end", cachePollController.EndPoll)
	}

	r.Run(PORT)
}
