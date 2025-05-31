package main

import (
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/", PersonHandler)
	http.HandleFunc("/", PerseHandler)

	log.Println("Starting server on :8081")
	log.Println("version 3")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
