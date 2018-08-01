package main

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type helloResponse struct {
	Value string `json:"value"`
}

func makeHealthEndpoint(svc HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := svc.Health()
		return nil, err
	}
}

func makeHealthHandler(svc HelloService) *httptransport.Server {
	return httptransport.NewServer(makeHealthEndpoint(svc), decodeRequest, encodeResponse)
}

func makeHelloEndpoint(svc HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.Hello()
		return &helloResponse{Value: v}, err
	}
}

func makeHelloHandler(svc HelloService) *httptransport.Server {
	return httptransport.NewServer(makeHelloEndpoint(svc), decodeRequest, encodeResponse)
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
