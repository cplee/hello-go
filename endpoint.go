package main

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type infoResponse struct {
	IPAddress string `json:"value"`
}

func makeHealthEndpoint(svc InfoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := svc.Health()
		return nil, err
	}
}

func makeHealthHandler(svc InfoService) *httptransport.Server {
	return httptransport.NewServer(makeHealthEndpoint(svc), decodeRequest, encodeResponse)
}

func makeInfoEndpoint(svc InfoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.Info()
		return &infoResponse{IPAddress: v}, err
	}
}

func makeInfoHandler(svc InfoService) *httptransport.Server {
	return httptransport.NewServer(makeInfoEndpoint(svc), decodeRequest, encodeResponse)
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if response == nil {
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}
