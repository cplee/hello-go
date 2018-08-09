package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Unexpected status %v", resp.Status)
	}
}
