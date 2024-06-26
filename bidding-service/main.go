package main

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	var svc BidderService
	svc = bidderService{}
	svc = loggingMiddleware{logger: logger, next: svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	bidderHandler := httptransport.NewServer(
		makeBidderEndpoint(svc),
		decodeBidderRequest,
		encodeResponse,
	)

	http.Handle("/", bidderHandler)
	http.Handle("/metrics", promhttp.Handler())
	err := logger.Log("msg", "HTTP", "addr", ":8081")
	if err != nil {
		return
	}
	err = logger.Log("err", http.ListenAndServe(":8081", nil))
	if err != nil {
		return
	}
}
