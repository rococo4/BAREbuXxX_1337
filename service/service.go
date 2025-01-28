package service

import (
	"barebuXxX_1337/meth"
	"barebuXxX_1337/service/Cache"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

type service struct {
	cacheUsers *Cache.Cache
	logger     *logrus.Logger
}

func New(cacheUsers *Cache.Cache, log *logrus.Logger) *service {
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
		s.logger.Info("Get request")
		w.Write([]byte("Got u"))
	})

	http.ListenAndServe(":8080", nil)
}
