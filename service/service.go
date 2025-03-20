package service

import (
	"barebuXxX_1337/logger"
	"barebuXxX_1337/meth"
	"barebuXxX_1337/service/Cache"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
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

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		meth.RequestCounter.Inc()
		start := time.Now()
		promhttp.Handler().ServeHTTP(w, r)
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues(r.Method, "200").Observe(duration)
	})
	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		meth.RequestCounter.Inc()
		start := time.Now()
		s.logger.Log("Got request", "info")
		w.Write([]byte("Got u"))
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues(r.Method, "200").Observe(duration)
	})

	http.ListenAndServe(":8080", nil)
}
