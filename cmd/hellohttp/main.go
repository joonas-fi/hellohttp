package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
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

	addr := ":80"

	fmt.Printf("hello world HTTP server listening at %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
