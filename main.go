package main

import (
	"barebuXxX_1337/meth"
	"barebuXxX_1337/service"
	"barebuXxX_1337/service/Cache"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(file)
	log.SetLevel(logrus.InfoLevel)
	cache := Cache.New(meth.CacheLen, log)
	service := service.New(cache, log)
	go cache.UpdateCache()
	service.Run()
}
