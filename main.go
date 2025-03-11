package main

import (
	"barebuXxX_1337/logger"
	"barebuXxX_1337/meth"
	"barebuXxX_1337/service"
	"barebuXxX_1337/service/Cache"
	"os"
)

func main() {
	lokiURL := os.Getenv("LOKI_URL")
	logger := logger.NewLokiLogger("barebuXxX_1337", 1, lokiURL)
	cache := Cache.New(meth.CacheLen, logger)
	service := service.New(cache, logger)
	go cache.UpdateCache()
	service.Run()
}
