package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-kit/kit/log"
)

// InfoResponse defines response structure
type InfoResponse struct {
	IPAddress    string    `json:"ipAddress"`
	SourceCommit string    `json:"sourceCommit"`
	Starttime    time.Time `json:"starttime"`
	Runtime      string    `json:"runtime"`
}

// InfoService provides information about service
type InfoService interface {
	Health() error
	Info() (*InfoResponse, error)
}

// infoService is a concrete implementation of InfoService
type infoService struct {
	logger    log.Logger
	startTime time.Time
}

// NewInfoService creates a new instance of InfoService
func NewInfoService() InfoService {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	svc := &infoService{
		logger:    logger,
		startTime: time.Now(),
	}
	return svc
}

func (s *infoService) Health() error {
	return nil
}

func (s *infoService) Info() (*InfoResponse, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	resp := &InfoResponse{
		SourceCommit: os.Getenv("SOURCE_COMMIT"),
		Starttime:    s.startTime,
		Runtime:      time.Now().Sub(s.startTime).String(),
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				resp.IPAddress = ipnet.IP.String()
				return resp, nil
			}
		}
	}

	return nil, fmt.Errorf("Unable to find ip address")
}
