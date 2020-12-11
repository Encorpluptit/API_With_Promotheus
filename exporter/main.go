package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math"
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
	seed float64
)

func createMetrics() {
	go func() {
		for {
			r := math.Sin(float64(time.Now().Unix())/10) + seed
			activity.Set(r)
			time.Sleep(1 * time.Second)
		}
	}()
}

func main() {
	s, _ := strconv.Atoi(os.Getenv("RAND_SEED"))
	seed = float64(s)
	createMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
