package Cache

import (
	"barebuXxX_1337/logger"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"strconv"
	"time"
)

type Cache struct {
	users          map[int64]string
	cacheLenMetric prometheus.Gauge
	logger         *logger.LokiLogger
}

func New(cacheLenMetric prometheus.Gauge, log *logger.LokiLogger) *Cache {
	return &Cache{
		users:          make(map[int64]string),
		cacheLenMetric: cacheLenMetric,
		logger:         log,
	}
}
func (c *Cache) getUsersAndWriteToCache() {
	size := rand.Intn(100) + 1
	c.cacheLenMetric.Set(float64(size))
	c.users = make(map[int64]string, size)
	for i := 0; i < size; i++ {
		c.users[int64(i)] = "user"
	}
}
func (c *Cache) UpdateCache() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		c.logger.Log("Update cache "+strconv.Itoa(len(c.users)), "info")
		c.getUsersAndWriteToCache()
	}
}
