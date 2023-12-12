package main

import (
	"log"
	"net/http"
	"time"

	"github.com/uber-go/tally"
	"github.com/uber-go/tally/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	reporter := prometheus.NewReporter(prometheus.Options{})
	
	scope, closer := tally.NewRootScope(tally.ScopeOptions{
		Reporter: reporter,
		// other options like prefix, tags, etc.
	}, 1*time.Second)
	defer closer.Close()

	counter := scope.Counter("my_counter")
	counter.Inc(1)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9090", nil))
}
