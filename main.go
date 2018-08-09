package main

import (
	"log"
	"net/http"
)

func main() {
	svc := NewInfoService()

	http.Handle("/health", makeHealthHandler(svc))
	http.Handle("/", makeInfoHandler(svc))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
