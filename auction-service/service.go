package main

type BidderService interface {
	FindAd(string) bidderResponse
}

type bidderService struct{}

func (bidderService) FindAd(s string) bidderResponse {
	adObject := AdObject{"", 0}
	return bidderResponse{s, adObject}
}

type ServiceMiddleware func(service BidderService) BidderService
