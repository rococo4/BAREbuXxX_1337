package meth

import (
	"github.com/prometheus/client_golang/prometheus"
)

var RequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Total number of HTTP requests.",
})
var CacheLen = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cache_length",
	Help: "The number of items in the cache.",
})
var HttpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_duration_seconds",
	Help:    "Duration of HTTP requests.",
	Buckets: prometheus.DefBuckets,
},
	[]string{"method", "status_code"})
