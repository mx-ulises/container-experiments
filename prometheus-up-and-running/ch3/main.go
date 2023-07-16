package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "puar_golang_http_request_count",
			Help: "Total number of HTTP requests received in this Golang App",
		},
	)

	requestErrorCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "puar_golang_http_request_error_count",
			Help: "Total number of HTTP requests errors received in this Golang App",
		},
	)

	coinCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "puar_golang_http_coin_count",
			Help: "Total number of coins reived in request",
		},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestErrorCount)
	prometheus.MustRegister(coinCount)
}

func requestCounterHandler(w http.ResponseWriter, r *http.Request) {
	// Increase request count
	requestCount.Inc()

	// Exception Handler
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprint(w, "Request Error: ", err)
			requestErrorCount.Inc()
		}
	}()

	// Evaluate and increase coin count
	queryParams := r.URL.Query()
	coinsString := queryParams.Get("coins")
	if coinsString == "" {
		coinsString = "1"
	}

	coins, err := strconv.ParseFloat(coinsString, 64)
	if err != nil {
		panic(err)
	}

	// Evaluate Random success rate at 20%
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Float64()
	if randomNumber < 0.2 {
		panic("Random Error")
	}

	// Increase number of coins
	coinCount.Add(coins)

	// Successful request
	fmt.Fprint(w, "Successful Request, added: ", coins)
}

func main() {
	http.HandleFunc("/counter", requestCounterHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
