package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", PersonHandler)

	log.Println("Starting server on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
