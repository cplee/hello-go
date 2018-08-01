package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/go-kit/kit/log"
)

// HelloService provides an API that hurts
type HelloService interface {
	Health() error
	Hello() (string, error)
}

// hellowService is a concrete implementation of HelloService
type helloService struct {
	logger log.Logger
	domain string
}

// NewHelloService creates a new instance of HelloService
func NewHelloService(domain string) HelloService {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	svc := &helloService{
		logger: logger,
		domain: domain,
	}
	return svc
}

func (s *helloService) Health() error {
	return nil
}

func (s *helloService) Hello() (string, error) {
	addrs, err := net.LookupHost(fmt.Sprintf("%s.%s", "hello-go", s.domain))
	if err != nil {
		return "", err
	}
	return strings.Join(addrs, ","), nil
}
