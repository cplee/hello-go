package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	domain, _ := os.LookupEnv("_SERVICE_DISCOVERY_NAME")
	svc := NewHelloService(domain)

	http.Handle("/health", makeHealthHandler(svc))
	http.Handle("/", makeHelloHandler(svc))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
