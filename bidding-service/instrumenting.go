package main

import (
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           BidderService
}

func (mw instrumentingMiddleware) FindAd(s string) (output bidderResponse) {
	defer func(begin time.Time) {
		lvs := []string{"method", "findAd", "error", ""}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output = mw.next.FindAd(s)
	return
}
