package service

import (
	"barebuXxX_1337/logger"
	"barebuXxX_1337/meth"
	"barebuXxX_1337/service/Cache"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type service struct {
	cacheUsers *Cache.Cache
	logger     *logger.LokiLogger
}

func New(cacheUsers *Cache.Cache, log *logger.LokiLogger) *service {
	return &service{
		cacheUsers: cacheUsers,
		logger:     log,
	}
}

func (s *service) Run() {
	meth.InitMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		meth.RequestCounter.Inc()
		s.logger.Log("Got request", "info")
		w.Write([]byte("Got u"))
	})

	http.ListenAndServe(":8080", nil)
}
