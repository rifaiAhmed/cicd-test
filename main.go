package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("listening on port 8081...")
	http.ListenAndServe(":8081", nil)
}
