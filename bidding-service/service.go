package main

import (
	"github.com/google/uuid"
	"math/rand/v2"
)

type BidderService interface {
	FindAd(string) bidderResponse
}

type bidderService struct{}

func (bidderService) FindAd(s string) bidderResponse {
	adObject := AdObject{uuid.New().String(), rand.IntN(100)}
	return bidderResponse{s, adObject}
}
