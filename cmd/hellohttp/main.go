package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/function61/gokit/net/http/httputils"
	"github.com/function61/gokit/os/osutil"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsServed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "requests_served",
			Help: "Count of served HTTP requests",
		},
	)
)

func main() {
	prometheus.MustRegister(requestsServed)

	osutil.ExitIfError(func(ctx context.Context) error {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hello world")

			requestsServed.Inc()
		})

		http.Handle("/metrics", promhttp.Handler())

		srv := &http.Server{
			Addr: ":80",
		}

		log.Printf("hello world HTTP server listening at %s", srv.Addr)

		return httputils.CancelableServer(ctx, srv, func() error { return srv.ListenAndServe() })
	}(osutil.CancelOnInterruptOrTerminate(nil)))
}
