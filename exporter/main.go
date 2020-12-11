package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	activity = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "activity",
		Help: "activity metric",
	})
)

func createMetrics() {
	go func() {
		for {
			r := rand.Float64() * 100
			activity.Set(r)
			time.Sleep(1 * time.Second)
		}
	}()
}

func main() {
	s, _ := strconv.Atoi(os.Getenv("RAND_SEED"))
	rand.Seed(int64(s))
	createMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
