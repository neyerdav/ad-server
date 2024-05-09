package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"io/ioutil"
	"net/http"
)

type bidderRequest struct {
	AdPlacementId string `json:"AdPlacementId"`
}

type AdObject struct {
	AdId     string `json:"AdId"`
	Bidprice int    `json:"bidprice"`
}

type bidderResponse struct {
	AdPlacementId string `json:"AdPlacementId"`
	AdObject      AdObject
}

func makeBidderEndpoint(svc BidderService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(bidderRequest)
		s := svc.FindAd(req.AdPlacementId)
		return s, nil
	}
}

func decodeBidderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request bidderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}
func decodeBidderResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response bidderResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
