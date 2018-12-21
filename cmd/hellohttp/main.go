package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("hello world\n")

	requestsServed := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "requests_served",
			Help: "Count of served HTTP requests",
		},
	)

	prometheus.MustRegister(requestsServed)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")

		requestsServed.Inc()
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
