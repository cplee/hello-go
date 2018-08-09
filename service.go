package main

import (
	"fmt"
	"net"
	"os"

	"github.com/go-kit/kit/log"
)

// InfoService provides an API that hurts
type InfoService interface {
	Health() error
	Info() (string, error)
}

// infoService is a concrete implementation of InfoService
type infoService struct {
	logger log.Logger
}

// NewInfoService creates a new instance of InfoService
func NewInfoService() InfoService {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	svc := &infoService{
		logger: logger,
	}
	return svc
}

func (s *infoService) Health() error {
	return nil
}

func (s *infoService) Info() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("Unable to find ip address")
}
