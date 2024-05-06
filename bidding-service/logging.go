package main

import (
	"github.com/go-kit/log"
	"time"
)

type loggingMiddleware struct {
	logger log.Logger
	next   BidderService
}

func (mw loggingMiddleware) FindAd(s string) (output bidderResponse) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "FindAd",
			"input", s,
			"took", time.Since(begin),
		)
	}(time.Now())
	output = mw.next.FindAd(s)
	return
}
